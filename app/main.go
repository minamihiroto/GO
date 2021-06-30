package main

import(
	"fmt"
)


func later() func(string)string {//func(string)stringが戻り値の型
	store := ""//無名関数の処理に関係する無名関数外の環境 👉 クロージャ

	return func(next string)string {//無名関数 👉 クロージャ
		s := store//ここでstoreを出すことで、無名関数の処理に関係し、storeはクロージャになる
		store = next//クロージャによって捕捉された変数の領域は、クロージャが何らかの形で参照され続ける限り値は破棄されない 👉 21行目から23行目まで参照され続けている
		return s
	}
}

func main(){
	f := later()//later()まで代入しているので、fを呼び出す時点で発火する
	n := later//nを呼び出し、引数を与えることで発火
	
	fmt.Println(f("AAA"))
	fmt.Println(f("BBB"))
	fmt.Println(f("CCC"))

	fmt.Println(n()("AAA"))//関数内の無名関数を呼び出す際は続けて引数をとる()を書く
	fmt.Println(n()("BBB"))
	fmt.Println(n()("CCC"))//nは関数later()を毎回新しく呼び直してるので、クロージャが新しくされて値が破棄されている
}
