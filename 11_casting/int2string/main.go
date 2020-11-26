package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "25"
	var result int
	var err error

	result, err = convert(s)
	fmt.Println("Result:", result)
	fmt.Println("Error:", err)
}

func convert(s string) (int, error) {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return r, nil
}
