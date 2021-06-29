package main

import(
	"fmt"
)

func main(){
	//interface{}は型名
	var x interface{}//全ての型と互換性がある

	fmt.Println(x)//<nil> 初期値にnilが入る

	x = 3.14
	fmt.Println(x)
	x = "文字列"
	fmt.Println(x)
	x = 1
	fmt.Println(x)

	//z := x+1 //interface{}はあくまで全ての型の値を汎用的に表す手段なだけであり、演算の対象としては使えない
}
