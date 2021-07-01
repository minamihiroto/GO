package main

import(
	m "math"//mでパッケージ名を上書き
	. "fmt"//.でパッケージ名を省略可
)

func main(){
	n := 1
	Println(n)
	//fmt.Println(n)//エラー
	Println(m.Pi)
	// Println(math.Pi)//エラー
}
