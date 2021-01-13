package main

// Service interface is responsible for managing the various services that will be used.
type Service interface {
	ListAll() ([]Book, error)
	Reserve(isbn string) (*Book, error)
	Return(isbn string) (*Book, error)
}

// BookService is a use case implementation of the basic Library functions
type BookService struct {
	repository Repository
}

func ValidateISBN(isbn string) error {
	return nil
}
