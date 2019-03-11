package models

// ModelError ...
type ModelError struct {
	Code int32 `json:"code"`

	Message string `json:"message"`

	Additional string `json:"additional,omitempty"`
}
