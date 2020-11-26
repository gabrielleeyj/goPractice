package main

import "fmt"

// Todo if declared as var it will not work.
// have to declare as a type.
type Todo map[int]string

// create interface for CRUD operations
type CRUD interface {
	create()
	remove()
	update()
	delete()
}

// Values store via a struct
type Values struct {
	Key   int
	Value string
}

type ValuesOutput Todo

func create(v Values) ValuesOutput {
	var out ValuesOutput

	out = map[int]string{v.Key: v.Value}

	return out
}

func test() {

	// must initialise via make()
	l := make(Todo)
	fmt.Printf("T[001] - %v\n", l)
	// declaring a map via [key]value
	l[1] = "Test this map."

	fmt.Printf("T[002] - %v\n", l)
	// fmt.Println(todo[1]) // returns Test this map.

	v := Values{Key: 1, Value: "Test"}
	create(v)
	fmt.Printf("T[003] - %v\n", l)

	v = Values{Key: 2, Value: "Test again"}
	create(v)
	fmt.Printf("T[004] - %v\n", l)

}

// func (v Values) create() Todo {
// 	return map[int]string{v.Key: v.Value}
// }
