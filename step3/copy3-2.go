package main

import "fmt"

func main() {
	//構造体の定義
	p := struct {
		name string
		age int
	}{name: "sanae", age: 20}

	//フィールドの呼び出し
	fmt.Println(p.name, p.age)  //sanae 20
}

