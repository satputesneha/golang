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

}
