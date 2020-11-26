package main

//6.1インターフェース　関数にインターフェースを埋め込む
//https://docs.google.com/presentation/d/1vKdJHHx4A_sP_Ft-3cakQqfJRbKxWSI9axzQ-GHHOoY/edit#slide=id.g4f15a7d687_0_214

/*
インターフェースの実用に近い例は、step3/pointer_reciever_sample.goを参照
*/

import "fmt"

//string型を返す関数型Funcを定義
type Func func() string

//レシーバがFunc型のf、メソッド名がString、返り値の型がstringのメソッドを定義
func (f Func) String() string {
	return f()
}

func main() {
	//fmt.stringerはfmtパッケージで定義されているインターフェース、右辺はクロージャに対してFunc型へのキャストを行っている
	var s fmt.Stringer = Func(func() string { return "hello" })
	fmt.Println(s) //hello
}
