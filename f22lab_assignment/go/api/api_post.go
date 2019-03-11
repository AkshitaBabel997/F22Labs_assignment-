package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "../../db"
	keyMgmt "../key_management"
	models "../models"
)

// PostCommentPost ...
func PostCommentPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var commentsInput models.CommentInput
	err := decoder.Decode(&commentsInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	} else {
		user, err := keyMgmt.GetValueFromKey(commentsInput.Apikey)
		if user == "" {
			errdata, _ := json.Marshal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errdata)
		} else {
			err = db.AddComment(commentsInput)
			fmt.Println(err)
			if err != nil {
				errdata, _ := json.Marshal(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(errdata)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Comment successfully added to the post"))
			}
		}
	}
}

// PostDislikePost ...
func PostDislikePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var postInput models.PostInput
	err := decoder.Decode(&postInput)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
	} else {
		user, err := keyMgmt.GetValueFromKey(postInput.Apikey)
		if err != nil || user == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error"))
		} else {
			err = db.DislikePost(postInput)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Error"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Post disliked successfully"))
			}
		}

	}
}

// PostLikePost ...
func PostLikePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var postInput models.PostInput
	err := decoder.Decode(&postInput)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
	} else {
		user, err := keyMgmt.GetValueFromKey(postInput.Apikey)
		if err != nil || user == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error"))
		} else {
			err = db.LikePost(postInput)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Error"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Post liked successfully"))
			}
		}

	}
}

// PostCreatePost ...
func PostCreatePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var createPost models.CreatePost
	err := decoder.Decode(&createPost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed creating post"))
	} else {
		// w.WriteHeader(http.StatusOK)
		user, err := keyMgmt.GetValueFromKey(createPost.Apikey)
		if err != nil || user == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Probablty you are not logged in"))
		} else {
			var postid = randStringRunes(32)
			var post = models.Posts{Id: postid, User: user, ImageURL: createPost.ImageURL, Likes: 0, Dislikes: 0, Content: createPost.Content}
			err = db.AddPost(post)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal error"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Post success Post ID: " + postid))
			}
		}
	}

}
