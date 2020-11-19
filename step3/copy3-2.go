package main

import "fmt"

/*
typeを使用しての構造体の定義
type user struct {
	name string
	age  int
}
*/

func main() {
	//構造体の定義 (typeを使用して型の定義を行うのが一般的)
	p := struct {
		name string
		age  int
	}{name: "sanae", age: 20}

	//フィールドの呼び出し
	fmt.Println(p.name, p.age) //sanae 20
}
