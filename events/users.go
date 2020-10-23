package main

import (
	"encoding/json"
	"fmt"
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
type Evaluate struct {
	Expr   string `json:"expr"`
	Result bool   `json:"result"`
}

func (app *RestApp) UserHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *RestApp) EvaluateExprHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("welcome evaluate expression")
	params := mux.Vars(r)
	v := params["expr"]
	fmt.Println(v)
	e := Evaluate{}
	e.Expr = v
	e.Result = checka(v)
	json.NewEncoder(w).Encode(&e)

}
func checka(photo string) bool {
	stack := make([]string, 0)
	m := make(map[string]string)
	m["("] = ")"
	m["{"] = "}"
	m["["] = "]"

	for i := 0; i < len(photo); i++ {

		person := string(photo[i])
		_, exists := m[person]
		if exists {
			stack = append(stack, person)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if person == m[top] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
		fmt.Println(person, stack)
	}
	return true
}
