package main

import (
	"fmt"
)

func main() {
	var n = 100
	// for loop from 0 to 100
	// if divisible by 3 return Fizz
	// if divisible by 5 return Buzz
	// if divisible by 3 and 5 return FizzBuzz
	for i := 0; i < n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
