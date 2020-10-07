package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}

func main() {
	router := mux.NewRouter()
	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})

	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	http.ListenAndServe(":8000", router)
}
