package main

import "fmt"

//奇数偶数の判定関数
func test(x int) (y int) {
	y = x % 2
	return
}

func main() {
	for i := 1; i <= 6; i++ {
		if test(i) == 0 {
			fmt.Printf("%d-偶数\n", i)
		} else {
			fmt.Printf("%d-奇数\n", i)
		}
	}
}
