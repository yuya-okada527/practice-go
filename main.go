package main

import "fmt"

func plus(x, y int) int {
	return x + y
}

var p = plus

func main() {
	// alias
	fmt.Println(p(1, 4))

}
