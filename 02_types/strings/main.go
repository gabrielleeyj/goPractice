package main

import "fmt"

func main() {
	var msg string
	msg = "Hello, World"

	fmt.Println(msg[0])

	fmt.Println(len(msg))

	newMsg := fmt.Sprintf("%s. We're ready!", msg)
	fmt.Println(newMsg)
}
