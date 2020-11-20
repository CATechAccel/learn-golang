package main

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
