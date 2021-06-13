package main

import "fmt"

func main() {
	// var n int = 100
	// var n = 200
	n := 80 // 省略形
	if (n >= 80){
		fmt.Println("80以上です。")
	} else if (n >= 60) {
		fmt.Println("60以上、80未満です。")
	} else {
		fmt.Println("60未満です。")
	}

	if !(n == 100){ // trueがfaleseになる
		fmt.Println("すごくない")
	}

	switch n % 2 {
	case 0:
		fmt.Println("偶数です。")
	case 1:
		fmt.Println("奇数です。")
	}
}
