package main

import (
	"fmt"
	"net/http"
)

//HTTPハンドラ（リクエストを処理する関数）の作成
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,HTTPサーバ")
}

func main() {
	//ハンドラとエントリポイント（パス）の結びつけ
	http.HandleFunc("/", handler) //第一引数の"/"がエントリポイント
	//HTTPサーバの起動
	http.ListenAndServe(":8080", nil)
}
