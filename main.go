package main

import (
	"fmt"
)

// パッケージ変数
var n = 10

// 型推論はできない
// n := 10

func main() {
	// 1行コメント
	var i int = 1
	fmt.Println(i)

	/*
		複行コメント
	*/
	var j, k int = 2, 3
	fmt.Println(j, k)

	var (
		a = 1
		b = true
	)

	fmt.Println(a, b)

	p := 1

	fmt.Println(p)

	fmt.Printf("n=%d\n", n+1)
}
