package main

//4.1型　型のキャスト
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_1095

import "fmt"

func main() {
	var f float64 = 10.1
	var n int = int(f)
	fmt.Println(f, n) //10.1 10
}
