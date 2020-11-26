package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("5:", function(5))
	fmt.Println("10:", function(10))
}

func function(x float64) float64 {
	return math.Pow(x, 2) - 5
}
