package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Book{})

	// Create
	db.Create(&Book{})

	// Read
	var book Book
	db.First(&book, 1)                 // find product with integer primary key
	db.First(&book, "code = ?", "D42") // find product with code D42

	// Update - update book's status
	db.Model(&book).Update("Status", "Reserved")

	// Update - update multiple fields
	db.Model(&book).Updates(Book{}) // non-zero fields
	db.Model(&book).Updates(map[string]interface{}{"Status": "Available"})

	// Delete - delete product
	db.Delete(&book, 1)
}
