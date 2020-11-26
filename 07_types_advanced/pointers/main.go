package main

import "fmt"

func main() {
	a := 42

	p := &a

	fmt.Println("Value:", a)
	fmt.Println("Pointer:", p)

	*p = 24

	fmt.Println("Value:", a)
	fmt.Println("Pointer:", p)

	A(a)
	fmt.Println("Call A:", a)

	B(&a)
	fmt.Println("Call B:", a)
}

func A(a int) {
	a = 25
}

func B(b *int) {
	*b = 25
}
