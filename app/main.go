package main

import(
	"fmt"
)

func init(){//main関数より先に実行される
	fmt.Println("init()")
}

func init(){//同じファイルに何回でも定義できる 👉 普通はしない
	fmt.Println("aaa")
}

func main() {
	fmt.Println("main()")
}
