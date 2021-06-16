package main

import (
	"fmt"
)

type Person struct{// 型に構造体を定義する、これをレシーバという
	Name string
	Age int
}

func main() {
	human := Person{}//Person型レシーバをhumanに定義
	human.Name = "minami"
	human.Age = 22
	fmt.Println(human)
	fmt.Println(human.Name)
	fmt.Println(human.Age)
}
