package main

import(
	"fmt"
)

func main() {
	//break ラベル
  LOOP:
	for{
		for{
			for{
				fmt.Println("開始")
				break LOOP// ラベル指定されてるfot文まで一気に抜ける
			}
			fmt.Println("ここは通らない")
		}
		fmt.Println("ここは通らない")
	}
	fmt.Println("完了")

	//continue ラベル
	L:
	for i := 1;i <=3; i++{
		for j:=1;j<=3; j++{
			if j > 1{//1超えたらここに入る
				continue L// ラベル指定されているfor文まで一気に戻る
			}
			fmt.Printf("%d*%d=%d\n",i,j,i*j)
	  }
	  fmt.Println("ここは通らない")
  }
}
