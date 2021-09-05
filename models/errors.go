package models

type CiscoError struct {
	Errors *Errors `json:"errors,omitempty"`
}

type Errors struct {
	Error []Error `json:"error,omitempty"`
}

type Error struct {
	ErrorMessage *string `json:"error-message,omitempty"`
	ErrorTag     *string `json:"error-tag,omitempty"`
	ErrorType    *string `json:"error-type,omitempty"`
}
