package main

//6.抽象化 【try】インターフェースを作ろう
//https://docs.google.com/presentation/d/1vKdJHHx4A_sP_Ft-3cakQqfJRbKxWSI9axzQ-GHHOoY/edit#slide=id.g4f15a7d687_0_267

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
