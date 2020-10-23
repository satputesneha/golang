package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type RestApp struct {
	db *sql.DB
	r  *mux.Router
}

func main() {
	fmt.Println("Welcome to rest app")

	app := &RestApp{}
	app.Initialise()

	log.Fatal(http.ListenAndServe(":8000", app.r))

	defer app.Teardown()
}
