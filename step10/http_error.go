package main

import (
	"net/http"
)

func hoge(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "エラーが発生したよ", 404)
}

func main() {
	http.HandleFunc("/", hoge)
	http.ListenAndServe(":8080", nil)
}
