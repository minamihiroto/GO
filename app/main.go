package main

import (
	"fmt"
)

type Result struct {//レシーバ
	X, Y int
}

func  Sum(num Result) int {//これはメソッドではなく、レシーバを引数に取ったただの関数
	return num.X + num.Y
}

func main() {
	I := Result{3, 4}//Iにレシーバを代入
	// fmt.Println(I.Sum())//この呼び出し方はできない
	fmt.Println(Sum(I))//関数を呼び出し、引数にレシーバを渡してる

	fmt.Println(I.X)
	fmt.Println(I.Y)
	fmt.Println(I.X + I.Y)//メソッドを呼び出した時と結果は同じ
}
