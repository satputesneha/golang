package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Event struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Host  string `json:"host"`
}

type restApp struct {
	db *sql.DB
	r  *mux.Router
}
type Users struct {
	User_id string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Age     int    `json:"age"`
}

func (app *restApp) EventHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rows, err := app.db.Query("select * from events where id=" + params["event_id"])
	e := Event{}
	if err != nil {
		json.NewEncoder(w).Encode(&e)
		return
	}
	rows.Next()

	rows.Scan(&e.Id, &e.Title, &e.Host)
	json.NewEncoder(w).Encode(&e)
}
func (app *restApp) UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome ")
	params := mux.Vars(r)
	rows, err := app.db.Query("select * from users where user_id=\"" + params["user_id"] + "\"")
	u := Users{}
	fmt.Println("one")
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(&u)
		return
	}
	rows.Next()
	fmt.Println("three")

	rows.Scan(&u.User_id, &u.Name, &u.Country, &u.Age)
	json.NewEncoder(w).Encode(&u)
}

func (app *restApp) Initialise() {
	var err error
	app.db, err = sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/events_db")
	if err != nil {
		panic(err)
	}
	app.initialiseHandlers()
}

func (app *restApp) initialiseHandlers() {
	app.r = mux.NewRouter()
	app.r.HandleFunc("/events/{event_id}", app.EventHandler)
	app.r.HandleFunc("/users/{user_id}", app.UserHandler)
}

func (app *restApp) Teardown() {
	app.db.Close()
}
