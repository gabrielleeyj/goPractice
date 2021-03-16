package db

import (
	"fmt"
	"testing"
)

func TestRepository(t *testing.T) {
	books := make([]Book, 0)
	repo := NewMemoryRepository(books)

	fmt.Println(repo)
}

func TestCreate(t *testing.T) {
	// initialize the memory repository
	books := make([]Book, 0)
	repo := NewMemoryRepository(books)

	fmt.Println(repo)

	book := Book{Name: "TestBook", ISBN: "1234567890"}
	// create function
	r, err := repo.Create(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestGetAll(t *testing.T) {
	// initialize the memory repository
	books := make([]Book, 0)
	repo := NewMemoryRepository(books)

	fmt.Println("T001: ", repo)

	book := Book{Name: "TestBook", ISBN: "1234567890"}
	// create function
	r, err := repo.Create(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("T002: ", r)

	// getall function
	err = repo.GetAll()
	if err != nil {
		fmt.Println(err)
	}
}
