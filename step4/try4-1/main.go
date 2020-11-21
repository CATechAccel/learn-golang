package main

import (
	"fmt"

	"./greeting"
)

func main() {
	var n string
	greeting.Do(&n)
	fmt.Println(n)
}
