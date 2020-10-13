package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to rest app")

	app := &restApp{}
	app.initialise()

	r := mux.NewRouter()
	r.HandleFunc("/events/{event_id}", app.eventHandler)
	log.Fatal(http.ListenAndServe(":8000", r))

	defer app.teardown()
}
