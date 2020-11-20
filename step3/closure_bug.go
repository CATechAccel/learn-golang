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
	/*fs[0],[1],[2]には、fmt.Println(i) という関数が入っているだけなので、
	10行目のforを利用して表示すると、7行目のforの処理終了時点のi(今回の場合２)が代入される。
	このため、0 1 2 ではなく、2 2 2という結果になる。
	（fsの長さを4とすると、3が4回表示される結果となる）
	*/
}
