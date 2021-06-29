package main

import(
	"fmt"
)

func plus(x,y int) int {
	return x+y
}

var plusAlias = plus//関数を代入、()をつけると実行してしまうのでつけない

func main(){
	fmt.Println(plus(10,5))
	fmt.Println(plusAlias(10,5))//別名で呼び出せる
}
