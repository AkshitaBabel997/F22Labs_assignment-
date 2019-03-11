package models

// User ...
type User struct {
	User string `json:"user" bson:"user"`

	Password string `json:"password" bson:"password"`
}
