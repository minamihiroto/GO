package main

import (
	"fmt"
)

func main() {
	name := "MINAMI"
	fmt.Println(&name)

	namePtr := &name// var namePtr *string = &nameと同義で、ポインタ型は「*変数の型」で定義できる
	fmt.Println(namePtr)//namePtrはnameのポインタを指している
	fmt.Println(*namePtr)//*をつけることで"ポインタが示す変数の値"と言う意味になる

	*namePtr = "HIROTO"//ポインタが示す変数の値に直接値を入れて更新することができる
	fmt.Println(name)
}
