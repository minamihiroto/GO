package main

import(
	"fmt"
)

func main(){
	//定数で判断
	n := 3
	switch n {
	case 1,2:
		fmt.Println("1or2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("unknown")
	}//3

	//fallthrough 👉 case処理を抜けない
	n = 1
	switch n {
	case 1,2:
		fmt.Println("1or2")
		fallthrough
	case 3:
		fmt.Println("3")
		fallthrough
	default:
		fmt.Println("unknown")
	}//1or2//3//unknown

	//簡易文 👉 switch分の中のみで有効な変数を定義
	switch 	n := 5; n {
	case 1,2:
		fmt.Println("1or2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("unknown")
	}//unknown

	//caseに式をかく 👉 定数が混在するとエラーが出る
	x:=4
	switch{
	case 0 < x && x < 3:
		fmt.Println("AAA")
	case 3 < x && n < 6:
		fmt.Println("BBB")
	}
	
}
