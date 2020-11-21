package main

//4.1型　配列
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_170

//4.1型　スライス
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_191

import "fmt"

func main() {
	//配列
	ns := [...]int{1, 2, 3, 4, 5}
	//要素にアクセス
	fmt.Println(ns[3]) //4
	//配列の長さの確認
	fmt.Println(len(ns)) //5
	//演算
	fmt.Println(ns[1:5]) //[2 3 4 5]

	//スライス → []のなかに要素数を何も書かないのが配列との違い！
	ns1 := []int{1, 2, 3, 4, 5}
	ns2 := make([]int, 3, 10) //長さ3,容量10のスライス（容量が埋まり始めたら随時追加される）
	//長さと容量の確認
	fmt.Println(len(ns2)) //3
	fmt.Println(cap(ns2)) //10
	//要素の追加
	ns1 = append(ns1, 6, 7)
	fmt.Println(len(ns1)) //7
	//要素の変更
	ns1[0] = 10
	fmt.Println(ns1[0:7]) //[10 2 3 4 5 6 7]
	//要素の削除
	ns1 = ns1[:1+copy(ns1[1:], ns1[3:])]
	fmt.Println(ns1[0:7]) //[10 4 5 6 7 6 7]
}
