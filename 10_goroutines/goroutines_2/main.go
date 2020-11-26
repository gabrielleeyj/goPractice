package main

import (
	"fmt"
	"time"
)

func main() {
	var sum int
	go func() {
		sum++
	}()
	fmt.Println(sum)
	time.Sleep(1 * time.Second)
	fmt.Println(sum)
}
