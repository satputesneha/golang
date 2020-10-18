package main

import "github.com/gorilla/mux"

func (app *restApp) initialiseHandlers() {
	app.r = mux.NewRouter()
	app.r.HandleFunc("/events/{event_id}", app.EventHandler)
	app.r.HandleFunc("/users/{user_id}", app.UserHandler)
	app.r.HandleFunc("/registration/{user_id}", app.RegHandler)
	app.r.HandleFunc("/allreg/{user_id}", app.AllRegHandler)
	app.r.HandleFunc("/allreg/byevent/{event_id}", app.AllRegByEventHandler)
}
