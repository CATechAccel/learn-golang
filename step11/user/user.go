package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"step11/db"
)

type userData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	/**
	QueryRowの使い方
	*/
	var u userData
	//QueryRow関数は、sql文で得られた結果を一件だけ取得する(SELECTなど)
	//`SELECT * FROM users`：usersテーブルから全件取得する
	row := db.DBInstance.QueryRow("SELECT * FROM users")
	//row.Scanでデータを構造体へとマッピングする
	if err := row.Scan(&u.ID, &u.Name, &u.Age); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}

func GetUserSlice(w http.ResponseWriter, r *http.Request) {
	/**
	Queryの使い方
	*/
	//Query関数は、sql文で得られた結果を全件取得する（SELECTなど）
	rows, err := db.DBInstance.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var userSlice []userData
	//rowsからの取り出しは、Next関数を利用する
	for rows.Next() {
		var u userData
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			panic(err)
		}
		userSlice = append(userSlice, u) //userSliceにuserDataを追加する処理
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(userslice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//ユーザーからのリクエストボディの中身を変数に格納
	body := r.Body
	defer body.Close()

	//bodyをbyte形式に変換
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		return
	}

	var reqBody userData
	//Unmarshal関数を使用すると、byteデータから任意の構造体にマッピングできる
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/**
	  Execの使い方
	*/
	//Execは、与えられたSQL文を実行する（INSERT,DELETE,UPDATEなど）
	//`INSERT INTO users(Name, Age) VALUES ('takuma', 22), ('hoge', 11)`：usersテーブルのNameとAgeカラムの値を指定して追加
	if result, err := db.DBInstance.Exec("INSERT INTO users(name, age) VALUES (?, ?)", reqBody.Name, reqBody.Age); err != nil {
		panic(err)
	} else {
		fmt.Println(result.LastInsertId()) //最後に追加されたデータのIDを取得する
		fmt.Println(result.RowsAffected()) //何件に影響したかを返す
	}
}
