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

	switch s {
	case 0:
		fmt.Println("凶")
	case 1, 2:
		fmt.Println("小吉")
	case 3, 4:
		fmt.Println("中吉")
	default:
		fmt.Println("大吉")
	}
}
