package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func yourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := fmt.Sprintf("<html><h1>Super Market</h1>hello Rest API %v\n</html>", vars)

	w.Write([]byte(s))
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("logging entrance")
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/products/{key1}", yourHandler).Methods("GET")
	r.HandleFunc("/category/{catname1}", yourHandler)
	r.Use(loggingMiddleware)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
