package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ToJSON encodes src to json *bytes.Buffer
func ToJSON(src interface{}) (*bytes.Buffer, string, error) {
	parsed, err := json.Marshal(src)
	if err != nil {
		return nil, "", err
	}
	return bytes.NewBuffer(parsed), "application/json", nil
}

var tag = "form"

type encoder func(reflect.Value) string

// ToForm encodes src to multipart/from-data *bytes.Buffer
func ToForm(src interface{}) (*bytes.Buffer, string, error) {
	v := reflect.ValueOf(src)
	res := new(bytes.Buffer)
	writer := multipart.NewWriter(res)
	defer writer.Close()
	err := toForm(v, writer)
	return res, writer.FormDataContentType(), err
}

func toForm(v reflect.Value, writer *multipart.Writer) error {

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return errors.New("form: interface must be a struct")
	}

	errors := Errors{}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		name, opts := fieldAlias(t.Field(i), tag)
		if name == "-" {
			continue
		}

		if reflect.TypeOf(v.Field(i).Elem()) == reflect.TypeOf((*os.File)(nil)).Elem() {
			writeFile(name, v.Field(i).Elem(), writer)
		}

		if isValidStructPointer(v.Field(i)) {
			writeJSON(name, v.Field(i).Elem().Interface(), writer)
			continue
		}

		if v.Field(i).Type().Kind() == reflect.Struct {
			writeJSON(name, v.Field(i).Elem().Interface(), writer)
			continue
		}

		enc := getEncoder(v.Field(i).Type())

		if enc != nil {
			value := enc(v.Field(i))
			if opts.Contains("omitempty") && isZero(v.Field(i)) {
				continue
			}
			_ = writer.WriteField(name, value)
		}

		if enc == nil {
			errors[v.Field(i).Type().String()] = fmt.Errorf("form: encoder not found for %v", v.Field(i))
			continue
		}
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}

func writeFile(name string, v reflect.Value, writer *multipart.Writer) error {
	file := reflect.ValueOf(v.Elem()).Interface().(os.File)
	fileContents, err := ioutil.ReadAll(&file)
	if err != nil {
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	part, err := writer.CreateFormFile(name, fi.Name())
	part.Write(fileContents)
	return nil
}

func writeJSON(name string, v interface{}, writer *multipart.Writer) {
	parsed, _ := json.Marshal(v)
	val := string(parsed)
	_ = writer.WriteField(name, val)
	return
}

type tagOptions []string

func (o tagOptions) Contains(option string) bool {
	for _, s := range o {
		if s == option {
			return true
		}
	}
	return false
}

func fieldAlias(field reflect.StructField, tagName string) (alias string, options tagOptions) {
	if tag := field.Tag.Get(tagName); tag != "" {
		alias, options = parseTag(tag)
	}
	if alias == "" {
		alias = field.Name
	}
	return alias, options
}

func parseTag(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}

func isValidStructPointer(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Ptr && v.Elem().IsValid() && v.Elem().Type().Kind() == reflect.Struct
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func:
	case reflect.Map, reflect.Slice:
		return v.IsNil() || v.Len() == 0
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}

func getEncoder(t reflect.Type) encoder {

	switch t.Kind() {
	case reflect.Bool:
		return boolEncoder
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intEncoder
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintEncoder
	case reflect.Float32:
		return float32Encoder
	case reflect.Float64:
		return float64Encoder
	case reflect.String:
		return stringEncoder
	case reflect.Ptr:
		f := getEncoder(t.Elem())
		return func(v reflect.Value) string {
			if v.IsNil() {
				return "null"
			}
			return f(v.Elem())
		}
	default:
		return nil
	}
}

func boolEncoder(v reflect.Value) string {
	return strconv.FormatBool(v.Bool())
}

func intEncoder(v reflect.Value) string {
	return strconv.FormatInt(int64(v.Int()), 10)
}

func uintEncoder(v reflect.Value) string {
	return strconv.FormatUint(uint64(v.Uint()), 10)
}

func floatEncoder(v reflect.Value, bits int) string {
	return strconv.FormatFloat(v.Float(), 'f', 6, bits)
}

func float32Encoder(v reflect.Value) string {
	return floatEncoder(v, 32)
}

func float64Encoder(v reflect.Value) string {
	return floatEncoder(v, 64)
}

func stringEncoder(v reflect.Value) string {
	return v.String()
}

// Errors contains information about errors.
type Errors map[string]error

func (e Errors) Error() string {
	s := ""
	for _, err := range e {
		s = err.Error()
		break
	}
	switch len(e) {
	case 0:
		return "(0 errors)"
	case 1:
		return s
	case 2:
		return s + " (and 1 other error)"
	}
	return fmt.Sprintf("%s (and %d other errors)", s, len(e)-1)
}
