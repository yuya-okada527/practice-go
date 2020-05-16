package main

import "fmt"

func main() {
	// lambda
	f := func(x, y int) int { return x + y }

	fmt.Println(f(1, 3))
}
