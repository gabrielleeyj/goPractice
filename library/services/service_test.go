package services

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"library.com/library/models"
	"library.com/library/repository"
	"testing"
)

func TestGetAll(t *testing.T) {
	// initialize repository
	database, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if database.Migrator().HasTable(&models.Book{}) {
		err = database.Migrator().DropTable(&models.Book{})
		if err != nil {
			t.Log("fail to drop table")
			t.FailNow()
		}
	}
	// create a book
	err = database.AutoMigrate(&models.Book{})
	if err != nil {
		t.Log("unable to initialize database", err)
		t.FailNow()
	}
	err = database.Create(&models.Book{Name: "The Lord of The Rings", ISBN: "9780544003415", Status: models.Available}).Error
	if err != nil {
		t.Log("unable to create book", err)
		t.FailNow()
	}
	r := repository.NewRepository(database)
	s := NewBookService(r)

	// get all
	books, err := s.ListAll()
	if err != nil {
		t.Log("fail to get all books", err)
		t.FailNow()
	}
	if len(books) != 1 {
		t.Log("fail to assert book length of 1")
		t.Fail()
	}
}

func TestReserve(t *testing.T) {
	// initialize repository
	database, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// create a book
	database.Create(&models.Book{Name: "The Lord of The Rings", ISBN: "9780544003415", Status: models.Available})

	r := repository.NewRepository(database)
	s := NewBookService(r)

	book, err := s.Reserve("9780544003415")
	if err != nil {
		panic("unable to reserve book")
	}
	fmt.Println(book)
}
