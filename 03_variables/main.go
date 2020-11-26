package main

import "fmt"

func main() {
	var x int
	var y int

	x = 5
	y = 3

	fmt.Println(x == y)

	y = y + 2

	fmt.Println(x == y)

	y++
	// y = y + 1

	fmt.Println(x == y)
}
