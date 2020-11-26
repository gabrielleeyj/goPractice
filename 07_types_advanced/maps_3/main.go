package main

import "fmt"

func main() {
	m := make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		m[fmt.Sprintf("Key%d", i)] = i
	}

	// fmt.Printf("Map: %+v\n", m)
	fmt.Println("Length:", len(m))

	doSomething(m, 33)

	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}

func doSomething(input map[string]int, value int) {
	input["AnotherMagicalValue!"] = value
}
