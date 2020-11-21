package main

import "fmt"

type T int //レシーバとできるのは、typeで定義した型/ポインタ型/ポインタを含む型
//tをレシーバとするメソッドfの定義
func (t *T) f() { fmt.Println("hi") }

func main() {
	var v T
	(&v).f() //v.f()でも良い
}
