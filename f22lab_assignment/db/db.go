package db

import (
	"context"
	"fmt"
	"time"

	models "../go/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

var mongoclient *mongo.Client
var database *mongo.Database
var databaseName = "test"
var userCollection = "user"
var postCollection = "post"
var commentCollection = "comment"

// StartMongo ...
func StartMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	mongoclient, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	err = mongoclient.Ping(ctx, readpref.Primary())

	defer cancel()

	if err != nil {
		panic(err)
	}

	database = mongoclient.Database(databaseName)

}

// AddUser ...
func AddUser(user models.User) error {

	collection := database.Collection(userCollection)
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

// GetUser ...
func GetUser(user string) (models.User, error) {
	var tempUser models.User
	collection := database.Collection(userCollection)
	err := collection.FindOne(context.Background(), bson.M{"user": user}).Decode(&tempUser)

	if err != nil {
		return models.User{}, err
	}
	return tempUser, nil
}

// AddPost ...
func AddPost(post models.Posts) error {
	collection := database.Collection(postCollection)
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// GetPost ...
func GetPost(postID string) (models.Posts, error) {
	var tempPost models.Posts
	collection := database.Collection(postCollection)
	err := collection.FindOne(context.Background(), bson.M{"_id": postID}).Decode(&tempPost)
	if err != nil {
		return models.Posts{}, err
	}
	return tempPost, nil
}

//LatestPosts ...
func LatestPosts() ([]models.Posts, error) {
	var posts []models.Posts
	var options options.FindOptions
	var limit int64 = 10
	options.Limit = &limit
	collection := database.Collection(postCollection)

	cur, err := collection.Find(context.Background(), bson.D{}, &options)
	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var elem models.Posts
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}
		posts = append(posts, elem)
	}
	return posts, nil
}

// UsersPosts ...
func UsersPosts(user string) ([]models.Posts, error) {
	var posts []models.Posts
	var options options.FindOptions
	var limit int64 = 10
	options.Limit = &limit
	collection := database.Collection(postCollection)

	cur, err := collection.Find(context.Background(), bson.M{"user": user}, &options)
	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var elem models.Posts
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}
		posts = append(posts, elem)
	}
	return posts, nil
}

// AddComment ...
func AddComment(comment models.CommentInput) error {
	_, err := GetPost(comment.Id)

	if err != nil {
		return err
	}

	fmt.Println(err)
	collection := database.Collection(commentCollection)
	_, err = collection.InsertOne(context.Background(), comment)

	if err != nil {
		return err
	}
	fmt.Println(err)
	return nil
}

// LikePost ...
func LikePost(postInput models.PostInput) error {
	var tempPost models.Posts
	collection := database.Collection(postCollection)
	err := collection.FindOne(context.Background(), bson.M{"_id": postInput.Id}).Decode(&tempPost)

	if err != nil {
		return err
	}
	updatedPost := tempPost
	updatedPost.Likes++
	err = collection.FindOneAndReplace(context.Background(), tempPost, updatedPost).Decode(&updatedPost)

	if err != nil {
		return err
	}
	return nil
}

// DislikePost ...
func DislikePost(postInput models.PostInput) error {
	var tempPost models.Posts
	collection := database.Collection(postCollection)
	err := collection.FindOne(context.Background(), bson.M{"_id": postInput.Id}).Decode(&tempPost)

	if err != nil {
		return err
	}
	updatedPost := tempPost
	updatedPost.Dislikes++
	err = collection.FindOneAndReplace(context.Background(), tempPost, updatedPost).Decode(&updatedPost)

	if err != nil {
		return err
	}
	return nil
}
