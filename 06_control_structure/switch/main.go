package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ADD = 1
	SUB = 2
	DIV = 3
	MTP = 4
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Calculator")
	fmt.Printf("\t%d) Add\n", ADD)
	fmt.Printf("\t%d) Substract\n", SUB)
	fmt.Printf("\t%d) Divide\n", DIV)
	fmt.Printf("\t%d) Multiply\n", MTP)

	fmt.Print("-> ")

	option, err := readIntegerInput(reader)
	if err != nil {
		panic(err)
	}

	fmt.Print("A -> ")
	a, err := readFloatInput(reader)
	if err != nil {
		panic(err)
	}

	fmt.Print("B -> ")
	b, err := readFloatInput(reader)
	if err != nil {
		panic(err)
	}

	if a > b {
		fmt.Println("A should be greater than B")
	}

	// TODO: Test difference between multiple if statements and switches with benchmarks.
	if option == ADD && a > b {
		fmt.Printf("\t%.2f + %.2f = %.2f\n", a, b, a+b)
	} else if option == SUB {
		fmt.Printf("\t%.2f - %.2f = %.2f\n", a, b, a-b)
	}

	fmt.Println("Result:")
	switch option {
	case ADD:
		fmt.Printf("\t%.2f + %.2f = %.2f\n", a, b, a+b)
	case SUB:
		fmt.Printf("\t%.2f - %.2f = %.2f\n", a, b, a-b)
	case DIV:
		fmt.Printf("\t%.2f / %.2f = %.2f\n", a, b, a/b)
	case MTP:
		fmt.Printf("\t%.2f * %.2f = %.2f\n", a, b, a*b)
	}
}

func readString(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error while reading input from console: %w", err)
	}

	output := strings.Replace(input, "\n", "", -1)

	return output, nil
}

func readIntegerInput(reader *bufio.Reader) (int, error) {
	input, err := readString(reader)
	if err != nil {
		return 0, err
	}

	output, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("error while converting input to integer: %w", err)
	}
	// This call ^
	// Error: error while converting input to integer
	// Base error: err

	// Next call (calling readIntegerInput)
	// Error: error while calling read integer input.
	// Base error: error while converting input to integer
	// Base error (2):

	return output, nil
}

func readFloatInput(reader *bufio.Reader) (float64, error) {
	input, err := readString(reader)
	if err != nil {
		return 0, err
	}

	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("error while converting input to float64: %w", err)
	}

	return output, nil
}
