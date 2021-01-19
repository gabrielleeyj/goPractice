package main

import (
	"errors"
)

// Repository interface is the implementation of the Repository functions.
type Repository interface {
	Get(isbn string) (*Books, error)
	GetAll() (*Books, error)
	Create(book Books) (*Books, error)
	Update(book Books) (*Books, error)
	Delete(isbn string) error
}

// BookRespository is the implementation of the database via Gorm
// type BookRespository struct {
// 	db *gorm.DB
// }

// func (r *BookRespository) Get(isbn string) (*Books, error) {
// 	return nil, errors.New("book not found")
// }

// func (r *BookRespository) GetAll(isbn string) (*Books, error) {
// 	return nil, errors.New("book not found")
// }

// func (r *BookRespository) Create(isbn string, name string) (*Books, error) {
// 	return nil, errors.New("book exists")
// }

// func (r *BookRespository) Update(isbn string, name string) (*Books, error) {
// 	return nil, errors.New("book cannot be updated")
// }

// func (r *BookRespository) Delete(isbn string) (*Books, error) {
// 	return nil, errors.New("book not found")
// }

// MemoryRepository is the implementation of data storage in memory.
type MemoryRepository struct {
	b []Books
}

// NewMemoryRepository stores an array of books
func NewMemoryRepository(initial []Books) Repository {
	return &MemoryRepository{
		b: initial,
	}
}

// Get searches for a single book in memory
func (m *MemoryRepository) Get(isbn string) (*Books, error) {
	for _, v := range m.b {
		if v.ISBN == isbn {
			return &Books{Name: v.Name}, nil
		}
	}
	return nil, errors.New("book not found")
}

// GetAll gets all the books in memory
func (m *MemoryRepository) GetAll() (*Books, error) {
	for _, v := range m.b {
		book := Books{ISBN: v.ISBN, Name: v.Name, Status: v.Status}
		return &book, nil
	}
	return nil, errors.New("book not found")
}

// Create allows the creation of a new book into memory.
func (m *MemoryRepository) Create(book Books) (*Books, error) {
	// checking function
	for _, v := range m.b {
		if v.ISBN == book.ISBN {
			return nil, errors.New("book exists")
		}
	}
	if ValidateISBN(book.ISBN) != false {
		m.b = append(m.b, book)
	}
	return &book, nil
}

// Update allows the updating of book name in the memory database.
func (m *MemoryRepository) Update(book Books) (*Books, error) {
	for _, v := range m.b {
		if v.ISBN == book.ISBN {
			book := Books{Name: book.Name}
			return &book, nil
		}
	}
	return nil, errors.New("book cannot be updated")
}

// Delete checks whether isbn exist and deletes
func (m *MemoryRepository) Delete(isbn string) error {
	temp := make([]Books, 0)
	for _, v := range m.b {
		if v.ISBN != isbn {
			temp = append(temp, v)
		}
	}

	return errors.New("book does not exist")
}
