package controller

import (
	"encoding/json"
	"net/http"

	"../entity"
	"../service"
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

var (
	postService service.PostService
)

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to get the posts"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (*controller) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to unmarshal the request"}`))
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid post"}`))
		return
	}

	result, err := postService.Create(&post)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
