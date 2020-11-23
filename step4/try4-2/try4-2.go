package main

//4.4ライブラリのバージョン管理　ライブラリを取得してみよう
//https://docs.google.com/presentation/d/13ykFLzbQXgiHebsPtmvsov5Kqi7R0OnNxUxY6BZZ_dA/edit#slide=id.g4f1426e3ae_0_248

import (
	"fmt"

	"github.com/tenntenn/greeting"
)

func main() {
	n := greeting.Do()
	fmt.Println(n)
}
