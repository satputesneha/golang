package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Event struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Host  string `json:"host"`
}

func (app *RestApp) EventHandler(w http.ResponseWriter, r *http.Request) {
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
