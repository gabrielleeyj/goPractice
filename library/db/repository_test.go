package db

import (
	"fmt"
	"library.com/library/models"
	"testing"
)

func TestRepository(t *testing.T) {
	books := make([]models.Book, 0)
	repo := NewMemoryRepository(books)

	fmt.Println(repo)
}

func TestCreate(t *testing.T) {
	// initialize the memory repository
	books := make([]models.Book, 0)
	repo := NewMemoryRepository(books)

	fmt.Println(repo)

	book := models.Book{Name: "TestBook", ISBN: "1234567890"}
	// create function
	r, err := repo.Create(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestGetAll(t *testing.T) {
	// initialize the memory repository
	books := make([]models.Book, 0)
	repo := NewMemoryRepository(books)

	fmt.Println("T001: ", repo)

	book := models.Book{Name: "TestBook", ISBN: "1234567890"}
	// create function
	r, err := repo.Create(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("T002: ", r)

	// GetAll function
	books, err = repo.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(books)
}
