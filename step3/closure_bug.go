package main

import "fmt"

func main() {
	fs := make([]func(), 3)
	for i := range fs {
		fs[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs {
		f()
	} //2 2 2

	/*
		fs[0],[1],[2]には、fmt.Println(i) という関数が入っているだけなので、
		10行目のforを利用してf()を呼び出した時には、iには７行目のfor分の処理を抜けて２が入っている。
		そして、クロージャは、関数のスコープ外の変数の値を参照するのではなく，内部的にはポインタを参照する。
		このため、0 1 2 ではなく、2 2 2という結果になる。
		（fsの長さを4とすると、3が4回表示される結果となる）
	*/
}
