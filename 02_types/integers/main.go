package main

import "fmt"

// Type		Size				Range
// int8		8 bits				-128 to 127
//	int16	16 bits				-215 to 215 -1
//	int32	32 bits				-231 to 231 -1
//	int64	64 bits				-263 to 263 -1
//	int		Platform dependent	Platform dependent

func main() {
	// Discrete -> 1, 2, 3, 4, 5 (int, int8, int16, int32, int64)
	// Continuous -> 1, 1.0000000000001, 1.00000000002, 1.00000000000003 (float32 float64)
	fmt.Println("1 + 1 =", 1+1)

	var valueB int64

	valueA := 1
	valueB = 225

	result := int64(valueA) + valueB

	fmt.Println(result)
}
