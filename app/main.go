package main

import(
	"fmt"
)

func main() {
	//goto文で同じ関数内の任意の位置（ラベル）までスキップする
	fmt.Println("A")
	goto L//Lまで処理を飛ばす
	fmt.Println("B")//実行されない
	// a := a //エラーが起きる 👉 変数定義は飛ばせない
	L://ラベル
	fmt.Println("C")
}
