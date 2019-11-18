// Package method implements all available telegram bot api methods.
package method

// TelegramMethod generic interface for telegram api methods.
type TelegramMethod interface {
	Method() string
	IsMultipart() bool
}

// SimpleMethod stub struct for simple methods.
type SimpleMethod struct{}

// IsMultipart returns true if method is multipart and false otherwise.
func (sm *SimpleMethod) IsMultipart() bool {
	return false
}

// MultipartMethod stub struct for multipart/form-data methods.
type MultipartMethod struct{}

// IsMultipart returns true if method is multipart and false otherwise.
func (mm *MultipartMethod) IsMultipart() bool {
	return true
}
