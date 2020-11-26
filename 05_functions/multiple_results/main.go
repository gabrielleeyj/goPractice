package main

import (
	"errors"
	"fmt"
)

// Main
// 1st Layer: runApplication
// swap(), doAction()

func main() {
	err := runApplication()
	if err != nil {
		panic(err)
	}
}

func swap(a, b string) (string, string) {
	return b, a
}

func runApplication() error {
	a, b := swap("world", "hello")
	fmt.Println(a, b)

	result, err := doAction(1)
	if err != nil {
		fmt.Println("There's been an error:", err)
		return nil
	}

	fmt.Println("Result:", result)
	return nil
}

func doAction(value int) (int, error) {
	return value * 10, errors.New("failed to do action")
}

// try and catch, finally.
