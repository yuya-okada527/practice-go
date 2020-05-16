package main

import (
	"fmt"
)

func main() {
	// 論理値
	a := false
	fmt.Println(a)

	// 10進数
	b := 10
	fmt.Println(b)

	// 8進数
	c := 0755
	fmt.Println(c)

	// 16進数
	d := 0x0718BEEF
	fmt.Println(d)

	// キャスト
	e := uint(17) // unsigned int
	fmt.Println(e)

	n := 1
	f := byte(n)
	g := int64(n)
	h := uint32(n)
	fmt.Println(f, g, h)

}
