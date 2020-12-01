package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Birth string `json:"birth"`
}

func getJson(w http.ResponseWriter, r *http.Request) {
	//ユーザーからのリクエストボディの中身を変数に格納
	body := r.Body
	defer body.Close()

	//bodyをbyte形式に変換
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		return
	}

	var reqBody person
	//Unmarshal関数を使用すると、byteデータから任意の構造体にマッピングできる
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reqBody.Name = "hoge"
	reqBody.Age = 1000
	reqBody.Birth = "fuga"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}

func main() {
	http.HandleFunc("/", getJson)
	http.ListenAndServe(":8080", nil)
}
