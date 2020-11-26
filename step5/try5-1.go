package main

import "fmt"

//インターフェースを定義
type stringer interface {
	String() string
}

type I string

func (i I) String() string {
	return "hello"
}

func new() stringer {
	return I("hoge")
}

func main() {
	var n I
	fmt.Println(n.String()) //hello
	fmt.Println(new())      //hello
}
