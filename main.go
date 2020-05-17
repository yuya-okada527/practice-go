package main

import "fmt"

func main() {
	// closure
	f := later()

	fmt.Println(f("Python"))
	fmt.Println(f("Java"))
	fmt.Println(f("JavaScript"))
	fmt.Println(f("Golang"))
	fmt.Println(f("Next?"))
}

func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}
