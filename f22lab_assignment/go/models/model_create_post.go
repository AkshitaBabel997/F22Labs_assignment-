package models

// CreatePost ...
type CreatePost struct {
	Content string `json:"content"`

	ImageURL string `json:"imageURL"`

	Apikey string `json:"apikey"`
}
