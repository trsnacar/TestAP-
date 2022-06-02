package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID       int    `json:"id`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

var dbUser []User

func main() {
	user := User{ID: 1, Email: "test@test.com", UserName: "cloud"}
	dbUser = append(dbUser, user)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, dbUser) //get post delete..
		} else {
			v, _ := ioutil.ReadAll(r.Body)
			fmt.Fprint(w, string(v))
			json.Unmarshal(v, &user)
		}
	})
	http.ListenAndServe(":8080", nil)
}
