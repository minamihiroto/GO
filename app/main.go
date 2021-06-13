package main

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().Unix())// 完全な乱数にするために必要
	fmt.Println(rand.Intn(10))// 上記がなければ実行する度同じ乱数
}
