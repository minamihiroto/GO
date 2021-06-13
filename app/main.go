package main

import "fmt"

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {//戻り値の個数分のデータ型を記述
	return y, x
}