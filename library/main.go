package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"library.com/library/models"
)

func main() {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Book{})

	// Create book table
	db.Create(&models.Book{Name: "The Lord of The Rings", ISBN: "9780544003415", Status: models.Available})
}

// // Read book entry
// var book Book
// db.First(&book, 1)                 // find product with integer primary key
// db.First(&book, "code = ?", "D42") // find product with code D42

// // Update - update book's status
// db.Model(&book).Update("Status", "Reserved")

// // Update - update multiple fields
// db.Model(&book).Updates(Book{}) // non-zero fields
// db.Model(&book).Updates(map[string]interface{}{"Status": "Available"})

// // Delete - delete product
// db.Delete(&book, 1)
