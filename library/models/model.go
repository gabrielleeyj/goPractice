package models

import "gorm.io/gorm"

// Status is the status of the book.
type Status string

var (
	// Reserved status
	Reserved Status = "Reserved"
	// Available status
	Available Status = "Available"
	// Unavailable status
	Unavailable Status = "Unavailable"
)

// Book is the model to store the books information
type Book struct {
	gorm.Model
	Name   string `json:"name"`
	ISBN   string `json:"isbn"`
	Status Status `json:"status"`
}
