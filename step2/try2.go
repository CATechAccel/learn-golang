package main

import "fmt"

func main() {
	//条件分岐にifを使用する場合
	for i := 1; i <= 4; i = i + 1 {
		if i%2 == 1 {
			fmt.Printf("%d - 奇数\n", i)
		} else {
			fmt.Printf("%d - 偶数\n", i)
		}
	}

	//条件分岐にswitchを使用する場合
	for i := 1; i <= 4; i += 1 {
		switch {
		case i%2 == 1:
			fmt.Printf("%d - 奇数\n", i)
		default:
			fmt.Printf("%d - 偶数\n", i)
		}
	}
}
