package main

import (
	"fmt"
)

type Result int//レシーバとしてtypeのみを定義、Resultというint型ができた（ただの別名）

func (num Result) Sum() int {//func後ろにレシーバを取るため、これはメソッド、numはResult型
	return  int(num)
}

func main() {
	I := Result(1)//Iにレシーバを代入
	fmt.Println(I.Sum())
}
