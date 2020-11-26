package main

import (
	"fmt"
)

func main() {
	var array [3]int

	fmt.Println("Size:", len(array))

	fmt.Println("Array [t=0]:", array)

	fmt.Printf("Array memory: %X\n", &array[0])
	fmt.Printf("Array memory: %X\n", &array[1])
	fmt.Printf("Array memory: %X\n", &array[2])

	for i := 0; i < len(array); i++ {
		array[i] = 3 * i
	}

	fmt.Println("Array [t=1]:", array)
}
