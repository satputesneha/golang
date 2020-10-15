package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func getObject(id1 string) Post {
	fmt.Println("hello world")
	db, err := sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/events_db")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected to db")
	}
	defer db.Close()
	rows, err := db.Query("select * from events where id=" + id1)

	rows.Next()
	p := Post{}
	err = rows.Scan(&p.ID, &p.Title, &p.Body)
	if err != nil {
		return Post{}
	}
	return p
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["sneha"])

	p := getObject(params["sneha"])

	json.NewEncoder(w).Encode(&p)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts/{sneha}", getPost).Methods("GET")

	http.ListenAndServe(":8000", router)
}
