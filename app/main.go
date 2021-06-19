package main

import "fmt"

func do(i interface{}) {//インターフェースをiに定義
	switch v := i.(type) {//i（型）をvに代入、型判定のswitch文
	case int:
		fmt.Printf("このvの型はint,%vの倍は%v\n", v, v*2)
	case string:
		fmt.Printf("このvの型はstring,%qの文字数は%v\n", v, len(v))
	default:
		fmt.Printf("このvの型は%T\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
