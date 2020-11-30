package main

//オリジナルのエラーの作成

import "fmt"

type originalError struct {
	point int
	cause string
}

func new(point int, cause string) error {
	return &originalError{
		point: point,
		cause: cause,
	}
}

//オリジナルのエラー関数を定義
func (o *originalError) Error() string {
	return fmt.Sprintf("%d エラーが起きました。エラー内容は%sです。", o.point, o.cause)
}

func main() {
	err := new(1, "hoge")
	if err != nil {
		panic(err.Error())
	}
}
