package main

import (
	"fmt"
	"strconv"
)

const message = `Calculator Application
========================
1) Add
2) Subtract
3) Divide
4) Multiply
=========================`

var option string
var a, b string
var result int

func main() {

	fmt.Println(message)
	fmt.Println("Select Option: ")
	fmt.Scanf("%v", &option)
	fmt.Println("Value 1: ")
	fmt.Scanf("%v", &a)
	fmt.Println("Value 2: ")
	fmt.Scanf("%v", &b)

	v1, _ := strconv.Atoi(a)
	v2, _ := strconv.Atoi(b)
	answer, _ := strconv.Atoi(option)

	switch answer {
	case 1:
		Add(v1, v2)
	case 2:
		Subtract(v1, v2)
	case 3:
		Divide(v1, v2)
	case 4:
		Multiply(v1, v2)
	}

}

// Add Function
func Add(a, b int) int {
	result := a + b
	fmt.Printf("Answer: %v", result)
	return result
}

// Subtract Function
func Subtract(a, b int) int {
	result := a - b
	fmt.Printf("Answer: %v", result)
	return result
}

// Divide Function
func Divide(a, b int) int {
	result := a / b
	fmt.Printf("Answer: %v", result)
	return result
}

// Multiply Function
func Multiply(a, b int) int {
	result := a * b
	fmt.Printf("Answer: %v", result)
	return result
}
