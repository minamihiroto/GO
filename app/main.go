package main

import(
	"fmt"
)

func main() {
	//配列の作成
	var a [10]int
	fmt.Println(a)
	//配列への代入
	a[0] = 1
	fmt.Println(a)
	fmt.Println(a[0])
	//配列の要素数を調べる
	fmt.Println(len(a))

	//スライスの作成
	s:=make([]int,10)
	fmt.Println(s)
	//スライスへの代入 👉 配列と同じ
	s[0] = 1
	fmt.Println(s)
	fmt.Println(s[0])
	//スライスの要素数を調べる 👉 配列と同じ
	fmt.Println(len(s))
}
