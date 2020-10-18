package main

import (
	"database/sql"
	"encoding/json"
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

type restApp struct {
	db *sql.DB
	r  *mux.Router
}
type Users struct {
	UserID  string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Age     int    `json:"age"`
}
type Registration struct {
	RegID   int    `json:"id"`
	EventID int    `json:"event_id"`
	UserID  string `json:"user_id"`
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
	log.Println("welcome UserHandler")
	params := mux.Vars(r)
	rows, err := app.db.Query("select user_id,Name,Country,Age  from users where user_id=?", params["user_id"])
	u := Users{}
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(&u)
		return
	}
	rows.Next()

	rows.Scan(&u.UserID, &u.Name, &u.Country, &u.Age)
	json.NewEncoder(w).Encode(&u)
}
func (app *restApp) RegHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("welcome to RegHandler")
	params := mux.Vars(r)
	rows, err := app.db.Query("select reg_id,event_id,user_id from registration where user_id =?", params["user_id"])
	reg := Registration{}
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(&reg)
		return
	}
	rows.Next()
	rows.Scan(&reg.RegID, &reg.EventID, &reg.UserID)
	json.NewEncoder(w).Encode(&reg)

}
func (app *restApp) AllRegByEventHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("welcome to AllRegByEventHandler")
	params := mux.Vars(r)
	rows, err := app.db.Query("select reg_id,event_id,user_id from registration where event_id =?", params["event_id"])
	reg := Registration{}
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(&reg)
		return
	}
	allregbyevent := make([]Registration, 0)
	for rows.Next() {

		rows.Scan(&reg.RegID, &reg.EventID, &reg.UserID)
		allregbyevent = append(allregbyevent, reg)

	}
	json.NewEncoder(w).Encode(&allregbyevent)
}

func (app *restApp) AllRegHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("welcome to AllRegHandler")
	params := mux.Vars(r)
	rows, err := app.db.Query("select reg_id,event_id,user_id from registration where user_id =?", params["user_id"])
	reg := Registration{}
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(&reg)
		return
	}
	allreg := make([]Registration, 0)

	for rows.Next() {
		rows.Scan(&reg.RegID, &reg.EventID, &reg.UserID)
		allreg = append(allreg, reg)

	}

	json.NewEncoder(w).Encode(&allreg)

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
	app.r.HandleFunc("/registration/{user_id}", app.RegHandler)
	app.r.HandleFunc("/allreg/{user_id}", app.AllRegHandler)
	app.r.HandleFunc("/allreg/byevent/{event_id}", app.AllRegByEventHandler)
}

func (app *restApp) Teardown() {
	app.db.Close()
}
