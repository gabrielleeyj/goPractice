package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	input := 25
	var result string

	result = convert(input)
	fmt.Println("Result:", result, reflect.TypeOf(result))
}

func convert(input int) string {
	return strconv.Itoa(input)
}
