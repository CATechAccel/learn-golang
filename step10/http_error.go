package main

import (
	"net/http"
)

func hoge(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, [エラーの詳細（string型）], [HTTPステータスコード(http.StatusXXXを利用)])
	http.Error(w, "エラーが発生したよ", http.StatusNotFound)
	//http.Error(w, "エラーが発生したよ", 404)と書くこともできる
}

func main() {
	http.HandleFunc("/", hoge)
	http.ListenAndServe(":8080", nil)
}
