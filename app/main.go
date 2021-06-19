package main

import "fmt"

func main() {
	var i interface{} = "hello"//空のインターフェースをiに定義、helloを代入

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// f = i.(float64) // panicエラー、中にあるのはstringだからアサーションが失敗したエラーが出る
	// fmt.Println(f)

	f, ok := i.(float64)
	fmt.Println(f, ok)//アサーションの成功失敗を第二引数で監視しているため、panicエラーが出ない

}