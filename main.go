package main

import "fmt"

func main() {
	// generator
	ints := integers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

}

func integers() func() int {
	var i int

	return func() int {
		i += 1
		return i
	}
}
