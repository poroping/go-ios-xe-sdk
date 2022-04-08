package models

type CiscoErrorResp struct {
	CiscoErrors *CiscoErrors `json:"errors,omitempty"`
}

type CiscoErrors struct {
	CiscoError []CiscoError `json:"error,omitempty"`
}

type CiscoError struct {
	ErrorMessage *string `json:"error-message,omitempty"`
	ErrorTag     *string `json:"error-tag,omitempty"`
	ErrorType    *string `json:"error-type,omitempty"`
	ErrorPath    *string `json:"error-path,omitempty"`
}
