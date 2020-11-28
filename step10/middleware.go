package main

import (
	"fmt"
	"net/http"
)

//ハンドラの前後で何かしらの処理を挟みたい時に使用する
func MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middlewareに入ったよ")
		next.ServeHTTP(w, r) //←ここがハンドラの呼び出し
		fmt.Println("ハンドラの処理が終わったよ")
	})
}

func do(w http.ResponseWriter, r *http.Request) {
	//rの中にはユーザーからのリクエスト情報が全て入っている
	//FormValueを使うと、クエリパラメータからキーを指定して値を取得できる
	msg := r.FormValue("msg")
	fmt.Fprintln(w, "hoge", msg)
}

func main() {
	//ハンドラをまとめるために使用
	mux := http.NewServeMux()
	mux.Handle("/middleware", MiddlewareFunc(http.HandlerFunc(do)))
	http.ListenAndServe(":8080", mux)
}
