package main

import "fmt"

func main() {//main関数が先に実行されるので関数の書く順番は関係ない
	totalAsk := ask("田中")
	fmt.Printf(totalAsk)
}

func ask(question string)string{
	return "こんにちは、"+question+"さん\n"
}