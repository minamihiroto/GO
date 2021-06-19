package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {//String() string、これもインターフェース(Stringer)
	//Stringerインターフェースはfmtパッケージに定義されている
	//fmtパッケージ以外にも変数を文字列で出力するためのインタフェースがある
	return fmt.Sprintf("%v (%v 歳)", p.Name, p.Age)
}

func main() {
	a := Person{"山田太郎", 42}
	z := Person{"田中エルフ", 9001}
	fmt.Println(a, z)
}
