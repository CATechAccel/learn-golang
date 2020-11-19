package main

import (
	"fmt"
)

func main() {
	//mapの初期化
	m := make(map[string]int) //m := map[string]int{"age": 20}でもできる
	//キーを指定して、要素の追加
	m["age"] = 20
	//キーを指定してアクセス
	fmt.Println(m["age"]) //20
	//存在を確認する
	n, ok := m["age"]
	fmt.Println(n, ok) //20 ture
	//キーを指定して削除
	delete(m, "age")
}
