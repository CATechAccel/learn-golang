package main

//4.1パッケージ　【try】パッケージを分けてみよう
//https://docs.google.com/presentation/d/13ykFLzbQXgiHebsPtmvsov5Kqi7R0OnNxUxY6BZZ_dA/edit#slide=id.g4f1426e3ae_0_235

import (
	"fmt"

	"./greeting"
)

func main() {
	var n string
	greeting.Do(&n)
	fmt.Println(n)
}
