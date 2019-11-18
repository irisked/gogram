package reply

// Button represents one button of the reply keyboard
// Optional fields are mutually exclusive.
type Button struct {
	Text           string
	MethodContact  bool `json:"method_contact,omitempty"`  // Optional
	MethodLocation bool `json:"method_location,omitempty"` // Optional
}

// NewButton creates a button
func NewButton(text string, options ...func(*Button)) *Button {
	button := new(Button)
	button.Text = text
	for _, option := range options {
		option(button)
	}
	return button
}

// Contact sets Contact to true and defaults everything other.
func Contact() func(*Button) {
	return func(btn *Button) {
		btn.MethodContact = true
		btn.MethodLocation = false
	}
}

// Location sets Location to true and defaults everything other.
func Location() func(*Button) {
	return func(btn *Button) {
		btn.MethodLocation = true
		btn.MethodContact = false
	}
}
