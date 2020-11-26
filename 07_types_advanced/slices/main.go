package main

import "fmt"

func main() {
	var slice []int

	fmt.Println("Size:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	fmt.Println("Size:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	anotherSlice := make([]int, 0, 500)

	fmt.Println("Size:", len(anotherSlice))
	fmt.Println("Capacity:", cap(anotherSlice))
}
