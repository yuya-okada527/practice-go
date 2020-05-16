package main

import "fmt"

func main() {
	// string型
	s := "string"
	fmt.Println(s)

	// raw string
	rs := `
	you can write multi line string like this!!
	Cool!!
	But this does not support tab alignment and escape for align.
	`
	fmt.Println(rs)

}
