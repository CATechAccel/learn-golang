package main

import "fmt"

//copy3-7でポインタを利用した参照渡しを行わなかった場合
func f2(x int) {
	fmt.Println(x) //5 ←コピーされた値であり、main関数内のxとは別物
	x = 100
	fmt.Println(x) //100
}

func main() {
	var x int = 5
	f2(x)
	fmt.Println(x) //5 ← コピーxが100に変更されても、main関数内のxには反映されない
}
