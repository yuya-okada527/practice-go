package main

import "fmt"

func main() {
	// function
	fmt.Println(plus(1, 2))
	hello()
	fmt.Println(divmod(5, 3))

	a, _ := divmod(17, 4)
	fmt.Println(a)
}

func plus(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello, Wolrd!")
}

func divmod(a, b int) (int, int) {
	return a / b, a % b
}
