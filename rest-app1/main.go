package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := fmt.Sprintf("<html><h1>Super Market</h1>hello Rest API %v\n</html>", vars)
	time.Sleep(10 * time.Second)
	w.Write([]byte(s))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/products/{key1}/{key2}", YourHandler)
	r.HandleFunc("/category/{catname1}", YourHandler)
	r.HandleFunc("/shopping-list/{shoplistnuum:[0-9]+}", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
