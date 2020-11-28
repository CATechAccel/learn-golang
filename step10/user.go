package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type userData struct {
	Name  string    `json:"name"`
	Age   int       `json:"age"`
	Birth time.Time `json:"birth"`
}

func handler2(response http.ResponseWriter, request *http.Request) {
	if _, err := fmt.Fprint(response, "HelloWorld"); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func getUser(response http.ResponseWriter, request *http.Request) {
	u := &userData{
		Name:  "紗苗",
		Age:   10,
		Birth: time.Now(),
	}

	//HTTPヘッダーの追加（json形式のデータであること,UTF-8の文字列であること）
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//HTTPステータスコードの追加（200:statusOK）
	response.WriteHeader(http.StatusOK)

	//userDataをjson形式に変換
	data, err := json.Marshal(u)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	//json化したbyteデータの送信
	_, _ = response.Write(data)
}

func createUser(response http.ResponseWriter, request *http.Request) {
	//ユーザーからのリクエストボディの中身を変数に格納
	body := request.Body
	defer body.Close()

	//bodyをbyte形式に変換
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		return
	}

	var reqBody userData
	//Unmarshal関数を使用すると、byteデータから任意の構造体にマッピングできる
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprint(response, "OK"); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/test", handler2)
	http.HandleFunc("/users", getUser)
	http.HandleFunc("/createUser", createUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("サーバの起動に失敗しました")
	}
}
