package main

import (
	"fmt"
)

func main() {
	name := "MINAMI"

	fmt.Println(name)
	changeName(&name)//ポインタを渡してる
	fmt.Println(name)

}

func changeName(namePtr *string) {//*をつけることで"ポインタが示す変数の値"と言う意味になるので、nameのポインタでなく値が入る型設定ができる（ポインタ型に指定）
	*namePtr = "HIROTO"//namePtrはポインタ型として型設定されてるため、ポインタ型変数として使える
}
