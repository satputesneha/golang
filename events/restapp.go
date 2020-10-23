package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func (app *RestApp) Initialise() {
	var err error
	app.db, err = sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/events_db")
	if err != nil {
		panic(err)
	}
	app.initialiseHandlers()
}

func (app *RestApp) Teardown() {
	app.db.Close()
}
