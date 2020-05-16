package main

import "fmt"

func main() {
	// operator
	a := "hoge"
	b := "fuga"
	fmt.Println(a + " " + b)

	fmt.Println(1 + 2)
	fmt.Println(2 * 3)
	fmt.Println(10 - 3)
	fmt.Println(12 / 5)
	fmt.Println(5 % 2)

	n := 1
	n += 2
	n *= 3
	n /= 4
	n %= 5
	fmt.Println(n)

	fmt.Println(1 == 1)
	fmt.Println(3 >= 2)
	fmt.Println(5 < 2)
	fmt.Println(4 != 2)
}
