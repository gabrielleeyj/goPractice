package main

import "fmt"

func main() {
	fmt.Println("AND:", and(true, true))
	fmt.Println("OR", or(true, false))
	fmt.Println("NOT:", not(true))
	fmt.Println("XOR", xor(true, false))
}

func and(a, b bool) bool {
	return a && b
}

func or(a, b bool) bool {
	return a || b
}

func not(value bool) bool {
	return !value
}

func xor(a, b bool) bool {
	return and(
		not(and(a, b)),
		or(a, b),
	)
}
