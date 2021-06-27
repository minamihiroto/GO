package main

import(
	"fmt"
	"gomodtest/animals"//moduleの名前/ディレクトリ名とする
)

func main(){
	fmt.Println(AppName())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}