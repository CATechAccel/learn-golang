package main

//4.1型　構造体リテラル
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_153

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
