package main

import "fmt"

func main() {
	fmt.Println(split(17))
}

func split(sum int) (x, y int) {//可読性が低いため、短い関数のみで使うべき
	x = sum - 5
	y = sum - x
	return//データ型定義部分にて何を返すかを書いてるからreturnだけでいい
}