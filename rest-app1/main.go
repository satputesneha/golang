package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := fmt.Sprintf("<html><h1>Super Market</h1>hello Rest API %v\n</html>", vars)

	w.Write([]byte(s))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/products/{key1}/{key2}", YourHandler).Methods("GET")
	r.HandleFunc("/category/{sneha}", YourHandler)
	r.HandleFunc("/shopping-list/{shoplistnuum:[0-9]+}", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
