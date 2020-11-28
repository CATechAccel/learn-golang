package main

//4.2関数　【try】ポインタ
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_362

import "fmt"

//ポインタを利用した値を入れ替える関数
func swap2(x, y *int) {
	*y, *x = *x, *y
	return
}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	fmt.Println(n, m) //20 10
}
