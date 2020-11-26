package main

import (
	"fmt"
	"time"
)

func main() {
	runWithClosure()

	time.Sleep(2 * time.Second)
}

func runWithClosure() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func runWithoutClosure() {
	for i := 0; i < 10; i++ {
		go printValue(i)
	}
}

func printValue(value int) {
	fmt.Println(value)
}
