package main

import(
	"fmt"
)

func main(){
	//範囲節によるfor 👉 配列をfor文で回す
	//for [配列のインデックス],[配列の要素] := range [配列型]{処理}
	fruits := [3]string{"Apple","Banana","Cherry"}
	for i,s := range fruits{
		fmt.Printf("fruits[%v]=%v\n",i,s)
	}
}
