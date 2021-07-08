package main

import(
	"fmt"
)

func main() {
	//panicは関数内の処理を中断させてdefer文のみを実行する
	//recoverはinterface{}型を返し、nil以外だった場合はpanicによる中断を回復させる
	defer func(){
		if x := recover(); x != nil{
			fmt.Println(x)
		}
	}()
	panic("panic!")
	fmt.Println("hello")//ここは実行されない
}
