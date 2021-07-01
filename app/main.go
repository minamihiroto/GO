package main

import(
	"fmt"
)

func main(){
	const(//定数定義の値省略
		X = 1
		Y
		Z
		R = "あ"
		L
		S = iota//constブロックのなかで定数が定義されるたびに1ずつ増える
		SS
		SSS
	)
	fmt.Println(X,Y,Z)
	fmt.Println(R,L)
	fmt.Println(S,SS,SSS)
}
