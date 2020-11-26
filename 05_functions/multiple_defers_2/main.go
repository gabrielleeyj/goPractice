package main

import "fmt"

func main() {
	defer fmt.Println("Hello 1")
	defer fmt.Println("Hello 2")
	defer fmt.Println("Hello 3")
	defer fmt.Println("Hello 4")
	defer fmt.Println("Hello 5")

	// Hello 1 -> Hello 2 -> ... -> Hello 5
	// Hello 5 -> Hello 4 -> ... -> Hello 1
}
