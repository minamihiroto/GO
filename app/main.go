package main

import "fmt"

func main() {//main関数が先に実行されるので関数の書く順番は関係ない
	ask(1,"山田")
	ask(2,"鈴木")
}

func ask(num int,question string){
	fmt.Println(num)
	fmt.Println("お名前を教えてください")
	fmt.Printf("こんにちは、%sさん\n",question)
}