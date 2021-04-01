package main

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"library.com/library/controllers"
	"library.com/library/models"
	"library.com/library/repository"
	"library.com/library/services"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		panic("Fail to automigrate")
	}
	// Create book table
	//repository.Create(&models.Book{Name: "The Lord of The Rings", ISBN: "9780544003415", Status: models.Available})

	r := repository.NewRepository(db)
	s := services.NewBookService(r)
	c := controllers.NewController(s)

	router := chi.NewRouter()
	router.Get("/", c.List)
	router.Put("/reserve/{isbn}", c.Reserve)
	router.Put("/return/{isbn}", c.Return)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic("Server down")
	}
}

// // Read book entry
// var book Book
// repository.First(&book, 1)                 // find product with integer primary key
// repository.First(&book, "code = ?", "D42") // find product with code D42

// // Update - update book's status
// repository.Model(&book).Update("Status", "Reserved")

// // Update - update multiple fields
// repository.Model(&book).Updates(Book{}) // non-zero fields
// repository.Model(&book).Updates(map[string]interface{}{"Status": "Available"})

// // Delete - delete product
// repository.Delete(&book, 1)
