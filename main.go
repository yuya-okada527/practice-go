package main

import "fmt"

func main() {
	// higher order function
	f := returnFunc()
	f()
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function.")
	}
}
