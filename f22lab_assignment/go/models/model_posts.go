package models

// Posts ...
type Posts struct {
	Id string `json:"_id" bson:"_id"`

	User string `json:"user" bson:"user"`

	ImageURL string `json:"imageURL" bson:"imageURL"`

	Content string `json:"content" bson:"content"`

	Likes int32 `json:"likes" bson:"likesd"`

	Dislikes int32 `json:"dislikes" bson:"dislikes"`

	Comments []Comments `json:"comments" bson:"comments"`
}
