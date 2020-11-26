package main

import (
	"fmt"
	"time"
)

func main() {
	go test() // go routine
	time.Sleep(50 * time.Microsecond)
}

func test() {
	fmt.Println("Test!")
}
