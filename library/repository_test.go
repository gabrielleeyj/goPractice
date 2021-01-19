package main

import (
	"fmt"
	"testing"
)

func TestRepository(t *testing.T) {
	books := make([]Books, 0)
	repo := NewMemoryRepository(books)

	fmt.Println(repo)
}

func TestCreate(t *testing.T) {
	// initialize the memory repository
	books := make([]Books, 0)
	repo := NewMemoryRepository(books)

	fmt.Println(repo)

	book := Books{Name: "TestBook", ISBN: "1234567890"}
	// create function
	r, err := repo.Create(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
