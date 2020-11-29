package main

import (
	"net/http"
	"step11/user"
)

func main() {
	http.HandleFunc("/user", user.GetUserData)
	http.HandleFunc("/users", user.GetUserSlice)
	http.HandleFunc("/createuser", user.CreateUser)
	http.ListenAndServe(":8080", nil)
}
