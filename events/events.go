package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Event struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Host  string `json:"host"`
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db, err := sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/events_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from events where id=" + params["event_id"])
	rows.Next()
	e := Event{}
	rows.Scan(&e.Id, &e.Title, &e.Host)
	json.NewEncoder(w).Encode(&e)
	w.Write([]byte(fmt.Sprintf("sneha %v", params)))
}

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/events_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/events/{event_id}", eventHandler)
	fmt.Println("hello world")
	log.Fatal(http.ListenAndServe(":8000", r))
}
