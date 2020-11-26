package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Status string

func main() {
	var value int
	value = 25
	print(value)

	result, err := castToInt(value)
	if err != nil {
		log.Println("Error while casting, error:", err)
	}
	print(result)

	var str string
	str = "Hello, world!"
	print(str)
	result, err = castToInt(str)
	if err != nil {
		log.Println("Error while casting, error:", err)
	}
	print(result)

	var status Status
	status = "Online"

	print(status)
}

// interface{}
// [25, int]
// ["Hello, world!", string]
// ["Online", main.Status]
func print(value interface{}) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type:", valueType)
	fmt.Println("Value:", value)
}

func castToInt(value interface{}) (int, error) {
	result, ok := value.(int)
	if !ok {
		return 0, errors.New("invalid data type when casting")
	}
	return result, nil
}
