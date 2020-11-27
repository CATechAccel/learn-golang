package main

import (
	"fmt"
	"net/http"
)

//p12の１
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,HTTPサーバ")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
