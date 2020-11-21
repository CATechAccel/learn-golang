package main

//4.2関数　【try】奇数偶数判定関数
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_348

import "fmt"

//奇数偶数の判定関数
func test(x int) (y int) {
	y = x % 2
	return
}

//返り値に変数名はつけず、型だけ指定するのが一般的（return ○○と明記することで可読性が上がる）
func test2(x int) int {
	return x % 2
}

func main() {
	for i := 1; i <= 6; i++ {
		if test(i) == 0 {
			fmt.Printf("%d-偶数\n", i)
		} else {
			fmt.Printf("%d-奇数\n", i)
		}
	}

	for i := 1; i <= 6; i++ {
		if test2(i) == 0 {
			fmt.Printf("%d-偶数\n", i)
		} else {
			fmt.Printf("%d-奇数\n", i)
		}
	}
}
