package main

import "fmt"

func main() {
	// Name: Score
	var studentScores map[string]int

	studentScores = make(map[string]int, 100)
	fmt.Println("Map:", studentScores)

	v, ok := studentScores["Marcos"]
	if !ok {
		fmt.Println("Marcos does not exist")
		return
	}

	studentScores["Marcos"] = 100

	v, ok = studentScores["Marcos"]
	if !ok {
		fmt.Println("Marcos does not exist")
		return
	}

	fmt.Println("Marcos does exist")
	fmt.Println("Marcos got:", v)

	delete(studentScores, "Marcos")

	v, ok = studentScores["Marcos"]
	if !ok {
		fmt.Println("Marcos does not exist")
	}

	fmt.Println("Marcos value has been deleted")
}
