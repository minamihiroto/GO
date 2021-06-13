package main

import "fmt"

func main() {
	// var n int = 100
	// var n = 200
	n := 300 // 省略形
	n = 400 // 変数の更新にはvarを意味する:をつけない
	n -= 100
	fmt.Println(n, n * 2)
}
