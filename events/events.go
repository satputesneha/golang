package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Welcome to rest app")

	app := &restApp{}
	app.Initialise()
	app.InitialiseHandlers()
	log.Fatal(http.ListenAndServe(":8000", app.r))

	defer app.Teardown()
}
