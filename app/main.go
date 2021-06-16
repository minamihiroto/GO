package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 10

	calculate(a,&b)//この処理でmain関数のaは10,bは11になってる

	fmt.Println("引数に整数を指定した場合：", a)
	fmt.Println("引数にポインタを指定した場合：", b)
}

func calculate(a int, b *int) {
	a += 1//あくまでここで定義されているaに更新が走ってるだけでmain関数のaとは関係ない
	*b += 1//main関数のbのポインタを参照してるポインタ型変数を更新してる
}
