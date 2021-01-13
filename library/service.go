package main

import "errors"

// Service interface is responsible for managing the various services that will be used.
type Service interface {
	ListAll() ([]Books, error)
	Reserve(isbn string) (*Books, error)
	Return(isbn string) (*Books, error)
}

// BookService is a use case implementation of the basic Library functions
type BookService struct {
	repository Repository
}

func ValidateISBN(isbn string) error {
	return nil
}

func (s *BookService) ListAll() ([]Books, error) {

	return nil, errors.New("No valid books")
}

func (s *BookService) Reserve(isbn string) (*Books, error) {
	return nil, errors.New("Unable to reserve book.")
}

func (s *BookService) Return(isbn string) (*Books, error) {
	return nil, errors.New("Unable to return book.")
}
