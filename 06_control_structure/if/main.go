package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := runAnotherExample(); err != nil {
		fmt.Println("Error:", err)
	}

	if err := runAnotherExample(); err != nil {
		fmt.Println("Second error:", err)
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d - even\n", i)
		} else {
			fmt.Printf("%d - odd\n", i)
		}
	}
}

func runAnotherExample() error {
	result, err := doAction(11)
	if err != nil {
		// Go 1.13
		return fmt.Errorf("error while running doAction with 1: %w", err)
	}

	fmt.Println("Result:", result)
	return nil
}

func doAction(value int) (int, error) {
	if value > 10 {
		return 0, errors.New("number bigger than 10")
	}
	return value * 10, nil
}
