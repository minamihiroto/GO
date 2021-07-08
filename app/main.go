package main

import(
	"fmt"
)

func main() {
	runDefer()
}

func runDefer(){
	//deferで定義すると関数の終了時に実行される
	defer fmt.Println("deferを実行2回目")
	defer fmt.Println("deferを実行")//最後に書かれたdeferがdeferの中で一番初めに実行される
	fmt.Println("done")
}
