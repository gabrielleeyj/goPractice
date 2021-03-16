package services

import (
	"fmt"
	"library.com/library/db"
	"library.com/library/models"
)

// Service interface is responsible for managing the various services that will be used.
type Service interface {
	ListAll() ([]models.Book, error)
	Reserve(isbn string) (*models.Book, error)
	Return(isbn string) (*models.Book, error)
}

// BookService is a use case implementation of the basic Library functions
type BookService struct {
	repository db.Repository
}

// ValidateISBN returns true if the provided string is a valid ISBN-10 or ISBN-13.
//
// The provided string must have a length of 10 or 13 and no formatting
// characters (spaces or hyphens).
//func ValidateISBN(isbn string) bool {
//	switch len(isbn) {
//	case 10:
//		return Validate10(isbn)
//	case 13:
//		return Validate13(isbn)
//	}
//	return false
//}
//
//// Validate10 returns true if the provided string is a valid ISBN-10.
////
//// The provided string must have a length of 10 and no formatting
//// characters (spaces or hyphens).
//func Validate10(isbn10 string) bool {
//	if len(isbn10) == 10 {
//		s, _ := sum10(isbn10)
//		return s%11 == 0
//	}
//	return false
//}
//
//// Validate13 returns true if the provided string is a valid ISBN-13.
////
//// The provided string must have a length of 13 and no formatting
//// characters (spaces or hyphens).
//func Validate13(isbn13 string) bool {
//	if len(isbn13) == 13 {
//		s, _ := sum13(isbn13)
//		return s%10 == 0
//	}
//	return false
//}
//
//// sum10 returns the weighted sum of the provided ISBN-10 string. It is used
//// to calculate the ISBN-10 check digit or to validate an ISBN-10.
////
//// The provided string must have a length of 9 or 10 and no formatting
//// characters (spaces or hyphens).
//func sum10(isbn string) (int, error) {
//	s := 0
//	w := 10
//	for k, v := range isbn {
//		if k == 9 && v == 88 {
//			// Handle "X" as the digit.
//			s += 10
//		} else {
//			n, err := strconv.Atoi(string(v))
//			if err != nil {
//				return -1, fmt.Errorf("Failed to convert ISBN-10 character to int: %s", string(v))
//			}
//			s += n * w
//		}
//		w--
//	}
//	return s, nil
//}
//
//// sum13 returns the weighted sum of the provided ISBN-13 string. It is used
//// to calculate the ISBN-13 check digit or to validate an ISBN-13.
////
//// The provided string must have a length of 12 or 13 and no formatting
//// characters (spaces or hyphens).
//func sum13(isbn string) (int, error) {
//	s := 0
//	w := 1
//	for _, v := range isbn {
//		n, err := strconv.Atoi(string(v))
//		if err != nil {
//			return -1, fmt.Errorf("Failed to convert ISBN-13 character to int: %s", string(v))
//		}
//		s += n * w
//		if w == 1 {
//			w = 3
//		} else {
//			w = 1
//		}
//	}
//	return s, nil
//}

// ListAll service functions.
func (s *BookService) ListAll() ([]models.Book, error) {
	books, err := s.repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("books not found %w", err)
	}
	return books, nil
}

// Reserve service functions
func (s *BookService) Reserve(isbn string) (*models.Book, error) {
	book, err := s.repository.Get(isbn)
	if err != nil {
		return nil, fmt.Errorf("book not found, %w", err)
	}
	book.Status = models.Reserved
	book, err = s.repository.Update(book.ID, *book)
	if err != nil {
		return nil, fmt.Errorf("failed to update %w", err)
	}
	return book, nil
}

// Return service functions
func (s *BookService) Return(isbn string) (*models.Book, error) {
	book, err := s.repository.Get(isbn)
	if err != nil {
		return nil, fmt.Errorf("book not found, %w", err)
	}
	book.Status = models.Available
	book, err = s.repository.Update(book.ID, *book)
	if err != nil {
		return nil, fmt.Errorf("failed to update %w", err)
	}
	return book, nil
}

func NewBookService(repository db.Repository) Service {
	return &BookService{
		repository: repository,
	}
}
