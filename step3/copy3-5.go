package main

//4.2関数　関数の定義
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_307

import "fmt"

//関数の定義:func 関数名(引数)(戻り値){}
func swap(x, y int) (x2, y2 int) {
	y2, x2 = x, y
	return
}

func main() {
	x, y := swap(10, 20)
	fmt.Println(x, y)
}
