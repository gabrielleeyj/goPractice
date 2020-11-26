package main

import (
	"fmt"
	"reflect"
)

type Result int

func main() {
	result := add(2, 3)
	fmt.Println("Value:", result)
	fmt.Println("TypeOf:", reflect.TypeOf(result))

	v := reflect.ValueOf(result)
	fmt.Println("ValueOf:", v)
	fmt.Println("Kind:", v.Kind())
}

func add(a, b int) interface{} {
	fmt.Println()
	return Result(a + b)
}
