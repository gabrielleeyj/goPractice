package main

import "fmt"

func generateCounter() func(int) int {
	var count int
	return func(value int) int {
		count += value
		return count
	}
}

func generateFib() func() int {
	x, y := 0, 1
	return func() (r int) {
		r = x
		x, y = y, x+y
		return
	}
}

func main() {
	counter := generateCounter()
	fib := generateFib()

	for i := 0; i < 10; i++ {
		fmt.Println("Count:", counter(2))
		fmt.Println("Fib:", fib())
	}
}
