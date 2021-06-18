package main

import (
	"fmt"
)

type Result struct{//構造体を持つレシーバ
	X, Y int
}

func (num Result) Sum() int {//メソッド
	return  num.X + num.Y
}

func (n *Result) Scale(f int) {//メソッドかつポインタを指定している
	n.Y = n.Y * f//2 * 3ということ、ここの2はIの変数定義の際に入っているがこれはポインタを指定してるので実現できてる
}

func main() {
	I := Result{1,2}//Iにレシーバを代入
	fmt.Println(I.Y)//2
	fmt.Println(I.X)//1
	fmt.Println(I.Sum())//1 + 2

	I.Scale(3)//Iの引き数（今回はY）を直接更新、3をscaleに渡してIにあるYを6に作り変えている
	fmt.Println(I.Y)//6
	fmt.Println(I.X)//1
	fmt.Println(I.Sum())//1 + 6
}
