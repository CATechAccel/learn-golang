package main

func main() {
	//条件分岐にifを使用する場合
	for i := 1; i <= 4; i = i+1{
		if i%2 == 1 {
			println("i - 奇数")
		} else {
			println("i - 偶数")
		}
	}

	//条件分岐にswitchを使用する場合
	for i:= 1; i <= 4; i += 1{
		switch {
		case i%2 == 1:
			println("i - 奇数")
		default:
			println("i - 偶数")
		}
	}
}
