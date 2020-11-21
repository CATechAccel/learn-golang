package main

//4.2関数　ポインタ
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_324

import "fmt"

func f(xp *int) {
	fmt.Println(xp) //0xc00001c080(main関数内で取得したxのアドレス)
	*xp = 100       //渡されたポインタ（アドレス）から中身にアクセスし、100を代入
}

func main() {
	var x int
	f(&x)          //xのポインタ（アドレス）を取得し、関数fに渡す
	fmt.Println(x) //100
}
