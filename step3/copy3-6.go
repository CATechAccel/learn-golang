package main

//4.2関数　関数型
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_821
//4.2関数　無名関数
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_816

import "fmt"

func print() string {
	return "hoge"
}

func main() {
	//配列の中に関数を要素として入れる
	fs := make([]func() string, 2)
	fs[0] = func() string { return "hoge" }
	fs[1] = func() string { return "fuga" }
	//for range スライス:第一返り値がindex番号、第２返り値が要素(value)
	for i, f := range fs {
		fmt.Println(i)
		fmt.Println(f())
	} //0 hoge 1 fuga

	m := make(map[string]string)
	m["name"] = "sanae"
	m["birth"] = "5月29日"
	//for range マップ:第一返り値がkey、第二返り値がvalue
	for key, value := range m {
		fmt.Println(key, value)
	} // name sanae birth 5月29日

	//無名関数:main関数の中で使用できる
	func() {
		fmt.Println("無名関数")
	}() //無名関数
}
