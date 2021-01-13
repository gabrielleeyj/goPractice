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
	db.AutoMigrate(&Books{})

	// Create book table
	db.Create(&Books{})

	// Read book entry
	var book Books
	db.First(&book, 1)                 // find product with integer primary key
	db.First(&book, "code = ?", "D42") // find product with code D42

	// Update - update book's status
	db.Model(&book).Update("Status", "Reserved")

	// Update - update multiple fields
	db.Model(&book).Updates(Books{}) // non-zero fields
	db.Model(&book).Updates(map[string]interface{}{"Status": "Available"})

	// Delete - delete product
	db.Delete(&book, 1)
}
