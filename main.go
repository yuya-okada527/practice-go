package main

import "fmt"

func main() {
	// higher order function
	callFunction(func() {
		fmt.Println("I'm a function.")
	})
}

func callFunction(f func()) {
	f()
}
