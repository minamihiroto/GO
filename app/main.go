package main

import(
	"fmt"
)

func main(){
	//初期値指定
	v := [3]int{1,2,3}//変数 := [要素数]要素の型{要素}
	n := [5]int{}
	var a [5]int
	z := [5]int{2,3,4}

	fmt.Println(v,n,a,z)

	//要素数の省略
	v1 := [...]int{1,2,3}
	v2 := [...]int{1,2,3,4,5}
	v3 := [...]int{}

	fmt.Println(v1,v2,v3)

	//代入
	v1[0] = 4
	v2 = [...]int{5,4,3,2,1}

	fmt.Println(v1,v2)
}
