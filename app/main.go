package main

import (
	"fmt"
)

func main() {
	Person.Name = "minami"
	Person.Age = 22
	fmt.Println(Person)
	fmt.Println(Person.Name)
	fmt.Println(Person.Age)
}

var Person struct{// 変数に構造体を定義する
	Name string
	Age int
}

// 構造体は非宣言文なので、何に（typeやvar）定義するのか決めないといけない
// struct {
//     Name string
//     Age  int
// }