package main

//4.3メソッド　【try】レシーバに変更を加える
//https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_406

import "fmt"

type Myint int

func (n *Myint) Inc() { *n++ } //ここでポインタ型にしておかないと、レシーバの変更は呼び出し元に反映されない

func main() {
	var n Myint = 1
	fmt.Println(n) //1
	n.Inc()        //nがポインタ型の時、(&n).Inc()と同じ意味
	fmt.Println(n) //2
}
