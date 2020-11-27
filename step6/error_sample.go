package main

//7.エラー処理
//https://docs.google.com/presentation/d/1QekRR0VE_5kEwUW7OatAx7OJmr65BHD5vkfFZ1lY13Y/edit#slide=id.g4fa5dcc83a_0_91

import (
	"errors"
	"fmt"
)

func main() {
	n, err := hello()
	//ifを利用してエラーの有無を調べる
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(n)

	//返り値がerror型のみの場合はこちらの書き方でif文を書く
	if err2 := sayHello(); err2 != nil {
		panic(err2.Error())
	}

}

func hello() (string, error) {
	return "", errors.New("エラーが起きました")
}

func sayHello() error {
	return errors.New("エラーが起きました") //errorsパッケージのNew関数
}
