package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type authMiddleware struct {
	passUsers map[string]string
}

func (a *authMiddleware) populate() {
	a.passUsers = make(map[string]string)
	a.passUsers["1234"] = "user1"
	a.passUsers["5678"] = "user2"

}
func (a *authMiddleware) MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := r.Header.Get("x-session-pass")
		user, ok := a.passUsers[pass]
		if ok {
			log.Printf("logged in already")
			log.Printf(user)
			next.ServeHTTP(w, r)
		} else {
			log.Printf("unauthorised user")
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})

}

func yourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := fmt.Sprintf("<html><h1>Super Market</h1>hello Rest API %v\n</html>", vars)

	w.Write([]byte(s))
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("logging entrance")
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/products/{key1}", yourHandler).Methods("GET")
	r.HandleFunc("/category/{catname1}", yourHandler)

	a := authMiddleware{}
	a.populate()
	r.Use(loggingMiddleware, a.MiddleWare)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
