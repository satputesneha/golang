package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Rest API\n"))
}
func AbcHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Rest API\n"))
}
func AbcdefHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("sneha %v", r.UserAgent())))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/Abcdef", AbcdefHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
