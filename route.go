package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{
		{
			Id:    1,
			Title: "Title1",
			Text:  "Text1",
		},
	}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	results, err := json.Marshal(&posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to marshal the posts"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to unmarshal the request"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)

	result, err := json.Marshal(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to marshal the post"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
