package main

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
}
