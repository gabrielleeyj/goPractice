package main

import (
	"fmt"
	"strings"
)

func slice() {
	// set capacity of slice to 5
	var slice [5]int
	// insert to slice
	slice[0] = 1
	fmt.Println("Index 0 of Slice", slice)

	// create empty slice
	var s []int
	fmt.Println("Empty Slice", s)
	// append 5 values to slice
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println("Append to empty slice", s)
	fmt.Println("Capacity of Slice: ", len(s))

	// append another 4 values to slice
	s = append(s, 6, 7, 8, 9)
	fmt.Println("Append to existing slice", s)
	// capacity of slice increases according to append
	fmt.Println("Capacity of Slice:", len(s))

	// if capacity of slice is 0, copy of slice has to be 0
	slice2 := make([]int, 0)
	fmt.Println("Make Slice", slice2)
	// if slice copy, is slice2[:5], returns slice bounds out of range with capacity 0
	slice3 := slice2[0:]
	fmt.Println("Copy Slice", slice3)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println("Length of Board", len(board))
	fmt.Println("Tic Tac Toe Board")

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}
