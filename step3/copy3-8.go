package main

//4.3メソッド　レシーバ
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_383

import "fmt"

type T int //レシーバとできるのは、typeで定義した型/ポインタ型/ポインタを含む型
//tをレシーバとするメソッドfの定義
func (t *T) f() { fmt.Println("hi") }

func main() {
	var v T
	(&v).f() //v.f()でも良い
}
