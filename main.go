package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Post struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Text   string `json:"Text"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", JobPostsHandler)

	http.ListenAndServe(":3333", r)
}
func JobPostsHandler(w http.ResponseWriter, r *http.Request) {

	posts := []Post{
		Post{"Post one", "John", "This is first post."},
		Post{"Post two", "Jane", "This is second post."},
		Post{"Post three", "John", "This is another post."},
	}

	json.NewEncoder(w).Encode(posts)
}
