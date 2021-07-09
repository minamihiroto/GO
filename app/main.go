package main

import(
	"fmt"
)

func sub(){
	for{//無限ループ
		fmt.Println(1)
  }
}

func sub2(){
	for{//無限ループ
		fmt.Println(2)
  }
}

func main() {
	go sub() //goで関数を実行すると、並列処理が行われる
	sub2()
}
