package apispec

// URL is the Bot API reference: the source of the spec.
const URL = "https://core.telegram.org/bots/api"

// Spec is the Bot API, as described by its reference page.
type Spec struct {
	Version     string             `json:"version"`      // e.g. "Bot API 10.3"
	ReleaseDate string             `json:"release_date"` // e.g. "August 24, 2026"
	Changelog   string             `json:"changelog"`    // link to the release notes
	Methods     map[string]*Method `json:"methods"`
	Types       map[string]*Type   `json:"types"`

	typeOrder []string // type names in page order
}

// Method is an API method, such as sendMessage.
type Method struct {
	Name        string   `json:"name"`
	Href        string   `json:"href"`
	Description []string `json:"description,omitempty"`
	Returns     []string `json:"returns"`
	Fields      []Field  `json:"fields,omitempty"`
}

// Type is an API object, such as Message. A type with Subtypes is a union
// of them, such as ChatMember.
type Type struct {
	Name        string   `json:"name"`
	Href        string   `json:"href"`
	Description []string `json:"description,omitempty"`
	Fields      []Field  `json:"fields,omitempty"`
	Subtypes    []string `json:"subtypes,omitempty"`
	SubtypeOf   []string `json:"subtype_of,omitempty"`
}

// Field is a parameter of a method or a field of a type. Types are spec
// names: "Integer", "String", "Array of PhotoSize", ...
type Field struct {
	Name        string   `json:"name"`
	Types       []string `json:"types"`
	Required    bool     `json:"required"`
	Description string   `json:"description"`
}
