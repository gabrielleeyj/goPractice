package main

import "fmt"

func main() {
	// AND
	fmt.Println(true && true)
	fmt.Println(true && false)

	// OR
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Print(false || false)

	// NOT
	fmt.Println(!true)
}
