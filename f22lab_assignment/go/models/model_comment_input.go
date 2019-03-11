package models

// CommentInput ...
type CommentInput struct {
	Apikey string `json:"apikey"`

	Id string `json:"_id,omitempty"`

	Comment string `json:"comment"`
}
