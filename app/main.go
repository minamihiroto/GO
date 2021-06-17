package main

import (
	"fmt"
)

type Result struct {//レシーバ
	X, Y int
}

func (num Result) Sum() int {//これがメソッド、メソッドはレシーバを関数名定義前に書く、numにResultを代入
	return num.X + num.Y
}

func main() {
	I := Result{3, 4}//Iにレシーバを代入することで、メソッドを呼び出すことができる
	fmt.Println(I.Sum())//足し算のメソッドを呼び出し

	fmt.Println(I.X)
	fmt.Println(I.Y)
	fmt.Println(I.X + I.Y)//メソッドを呼び出した時と結果は同じ
}
