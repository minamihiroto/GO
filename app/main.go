package main

import "fmt"

func main() {
	var input string// input := 👈 これではエラーが出てしまう
	fmt.Println("お名前を教えてください")
	fmt.Scan(&input)
	fmt.Printf("こんにちは、%sさん",input)
}
