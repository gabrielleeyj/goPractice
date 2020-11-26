package main

import (
	"bufio"
	"fmt"
	"os"
)

const msg = `ToDo List ----
1. Create Todo
2. Read List
3. Update Todo
4. Delete Todo
----------------
Select Option: `

const dmsg = `Delete Item No: `

// TodoList to store the data
type TodoList map[int]string

func input(message string) (answer int) {
	fmt.Println(message)
	fmt.Scanln(&answer)
	return
}

// readList show all the todos
func readList(list TodoList) {
	for k, v := range list {
		fmt.Printf("ToDo - %v : %v\n", k, v)
	}
	return
}

func main() {

	// deprecated value string variable replaced with bufio line
	// var value string

	// initialise todolist
	td := make(TodoList)
	td[0] = "test a string"
	fmt.Printf("T[001] -- %v\n", td)

	// fmt package intentionally removes white space thus should use bufio package
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("")
		fmt.Printf("Current ToDos: %v\n", len(td))
		Option := input(msg)

		switch Option {
		case 1:
			i := len(td)
			if td[i] != td[0] {
				fmt.Println("Input Todo Text: ")
				// fmt package intentionally removes white space thus should use bufio package
				// thus removed fmt.Scanln(&value)
				if scanner.Scan() {
					line := scanner.Text()
					fmt.Printf("Input was: %q\n", line)
					td[i] = line
					i++
				} else {
					fmt.Println("Input Todo: ")
					// fmt package intentionally removes white space thus should use bufio package
					// thus removed fmt.Scanln(&value)
					if scanner.Scan() {
						line := scanner.Text()
						fmt.Printf("Input was: %q\n", line)
						td[i] = line
						i++
					}

				}
			}
		case 2:
			readList(td)
		case 3:
			var input int
			fmt.Println("Input Todo No: ")
			fmt.Scanln(&input)
			fmt.Printf("Editing Todo No: %v\n", input)

			if scanner.Scan() {
				line := scanner.Text()
				fmt.Printf("Input was: %q\n", line)
				td[input] = line
			}

		case 4:
			d := input(dmsg)
			delete(td, d)
		default:
			fmt.Println("Invalid Option")
		}

	}
}
