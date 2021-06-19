package main

import "fmt"

type I interface {//インターフェースは中に関数が持てる、typeなので型
	M()
}

type T struct {//ストラクトは中に変数が持てる、typeなので型
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}//T型のhelloがiに入っていて、iはI型
	i.M()//iはインターフェースのI型なので、M関数が使える
}
