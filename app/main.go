package main

import "fmt"

func main() {
	var i I
	var t *T//*は"ポインタが示す変数の値"、ここで値はnil
	i = t//I型のiにT型のtが入っただけ
	P(i)//<nil>, *main.T
	i.M()//<nil>

	i = &T{"hello"}//T型のポインタをとってきてhelloをiに渡している、ポインタ先に値helloを指定
	P(i)//&{hello}, *main.T
	i.M()//hello
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {//メソッドにポインタを渡している
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func P(i I) {
	fmt.Printf("%v, %T\n", i, i)//%vは値 == nil、%Tは型 == *main.T
}
