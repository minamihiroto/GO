package main

import(
	"fmt"
)

func main(){
	//if 条件式{処理}
	//if 簡易文; 条件式{処理}
	if x,y := 1,2; x < y{//x,y := 1,2;の部分が簡易文
		fmt.Println(x,y)
	}

	//エラーが発生した時にも使える
	//if _,err := 何らかの関数(); err != nil{
	//	エラーが発生したときの処理
	//}

	// for 条件式{処理}
	// for 初期化文; 条件式; 後処理文{処理}//古典的な書き方
	
	//break文は処理を中断させる (一番近くのfor,switch,selectを抜ける)
	//continue文は残っている処理をスキップして次のループ処理にいく(一番近くのforで使用)

	i := 0
	// i := 10//breakで抜けない
	for i<30{
		fmt.Println(i)
		i++

		if i == 10{
			fmt.Printf("%vになったから出るよ\n",i)	
			break
		}
	}

	for n:=0;n<10;n++{
		if(n%2==1){
			continue
		}
		fmt.Println(n)
		n++
	}


}
