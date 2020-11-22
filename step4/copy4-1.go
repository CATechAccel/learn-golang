package main

import "fmt"

//4.2パッケージ変数とスコープ　パッケージ変数
//https://docs.google.com/presentation/d/13ykFLzbQXgiHebsPtmvsov5Kqi7R0OnNxUxY6BZZ_dA/edit#slide=id.g4f1426e3ae_0_631

//パッケージ変数 → 関数外で定義した変数は、パッケージ内であればどこでも利用できる（この時、:=では定義できない）
var msg string = "hello"

func f() { fmt.Println(msg) }

func main() {
	f() //hello
	msg = "world"
	f() //world
}
