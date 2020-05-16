package main

import "fmt"

func main() {
	// 配列
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)

	// 参照
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[3])
	fmt.Println(a[4])

	// 初期値
	var b [5]int
	fmt.Println(b)

	// 要素数の省略
	c := [...]int{1, 2, 3}
	fmt.Println(c)

	// 要素の代入
	c[1] = 0
	fmt.Println(c)
}
