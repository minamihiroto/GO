package main

import (
	"fmt"
)

func main() {
	score := 0
	plusTen(score)//scoreと値は同じ0を関数に渡しているが、コピーなのでポインタは違うもの

	fmt.Println(score)//0が表示される
	fmt.Println(&score)//0のポインタ
}

func plusTen(score int) {
	score += 10//0に10をたす

	fmt.Println(score)//10が表示される
	fmt.Println(&score)//10のポインタ
}
