package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func getAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []post // Create a new slice for easier encoding
	for _, p := range postsMap {
		posts = append(posts, p)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	post, exists := postsMap[id]

	if !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	newPost, err := createPostHandlerInternal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}

func createPostHandlerInternal(r *http.Request) (post, error) {
	var newPost post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newPost); err != nil {
		log.Fatal(err)
		return post{}, err
	}
	newPost.ID = uuid.New()
	newPost.Datetime = time.Now()
	postsMap[newPost.ID.String()] = newPost
	return newPost, nil
}