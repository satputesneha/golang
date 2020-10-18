package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Users struct {
	UserID  string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Age     int    `json:"age"`
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
