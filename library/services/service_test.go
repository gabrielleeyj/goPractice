package services

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"library.com/library/db"
	"library.com/library/models"
	"testing"
)

func TestGetAll(t *testing.T) {
	// initialize db
	database, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// create a book
	database.Create(&models.Book{Name: "The Lord of The Rings", ISBN: "9780544003415", Status: models.Available})

	r := db.NewRepository(database)
	s := NewBookService(r)

	// get all
	books, err := s.ListAll()
	if err != nil {
		panic("no books")
	}
	fmt.Println(books)

}
