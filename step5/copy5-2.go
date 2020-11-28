package main

//6抽象化　埋め込みとフィールド
//https://docs.google.com/presentation/d/1vKdJHHx4A_sP_Ft-3cakQqfJRbKxWSI9axzQ-GHHOoY/edit#slide=id.g4f15a7d687_0_284

import "fmt"

type Hoge struct {
	N int
}

//埋め込み：型リテラル（配列型や構造体型、スライス型など）以外のtypeで定義した型や組み込み型、インターフェースは埋め込める
type Fuga struct {
	Hoge
}

func main() {
	f := Fuga{Hoge{N: 100}}
	g := Hoge{N: 200}

	fmt.Println(f.N)      //100
	fmt.Println(f.Hoge.N) //100（型名を指定してアクセスすることもできる）
	//埋め込みは継承ではないが、継承と同じようなことができる
	fmt.Println(g.N) //200
}
