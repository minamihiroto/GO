package main

import(
	"fmt"
)

func returnFunc() func() {//func()が型
	return func() {
		fmt.Println("これは関数処理の結果を返しています")
	}
}

func main(){
	//無名関数の基本
	nanashi := func(x,y int)int{return x+y}//変数に関数リテラルを入れている
	fmt.Println(nanashi(2,3))//変数()で呼び出せる

	//関数を直接呼び出すこともできる
	fmt.Printf("%#v\n",func(x,y int)int{return x+y}(2,3))

	//関数を返す関数
	f := returnFunc()
	f()
}
