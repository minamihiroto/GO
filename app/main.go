package main

import(
	"fmt"
)


func div(a,b int)(int,int){
	q := a+b
	r := a-b
	return q,r
}

func main(){
	Q,R := div(2,1)	//割り振り代入
	fmt.Println(Q,R)

	q1,_ := div(3,4)//戻り値の破棄
	fmt.Println(q1)

}
