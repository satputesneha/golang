package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
