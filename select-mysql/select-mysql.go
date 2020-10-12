package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("hello world")
	db, err := sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/events_db")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected to db")
	}
	defer db.Close()
	rows, err := db.Query("select * from events ")
	var id int
	var name, host string
	for rows.Next() {
		err = rows.Scan(&id, &name, &host)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, name, host)
	}

	rows, err = db.Query("replace into events_db.events(id, Name, Host) values(5,'Hourly_meeting','David Johnson')")
	if err != nil {
		fmt.Println(err)
	}
}
