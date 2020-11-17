package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//現在時刻をナノ秒で表し変数に代入
	t := time.Now().UnixNano()
	//現在時刻を乱数の種として設定
	rand.Seed(t)
	s := rand.Intn(6) /*0-5の乱数*/

	switch s + 1 {
	case 1:
		fmt.Println("凶")
	case 2, 3:
		fmt.Println("小吉")
	case 4,5:
		fmt.Println("中吉")
	default:
		fmt.Println("大吉")
	}
}
