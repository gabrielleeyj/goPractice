package main

import (
	"errors"

	"gorm.io/gorm"
)

// Repository interface is the implementation of the Repository functions.
type Repository interface {
	Get(isbn string) (*Books, error)
	GetAll(isbn string) (*Books, error)
	Create(isbn, name string) (*Books, error)
	Update(isbn, name string) (*Books, error)
	Delete(isbn string) error
}

// BookRespository is the implementation of the database via Gorm
type BookRespository struct {
	db *gorm.DB
}

func (r *BookRespository) Get(isbn string) (*Books, error) {
	return nil, errors.New("book not found")
}

func (r *BookRespository) GetAll(isbn string) (*Books, error) {
	return nil, errors.New("book not found")
}

func (r *BookRespository) Create(isbn string, name string) (*Books, error) {
	return nil, errors.New("book exists")
}

func (r *BookRespository) Update(isbn string, name string) (*Books, error) {
	return nil, errors.New("book cannot be updated")
}

func (r *BookRespository) Delete(isbn string) (*Books, error) {
	return nil, errors.New("book not found")
}

// MemoryRepository is the implementation of data storage in memory.
type MemoryRepository struct {
	b []Books
}

func (m *MemoryRepository) Get(isbn string) (*Books, error) {
	for _, v := range m.b {
		if v.ISBN == isbn {
			return &Books{Name: v.Name}, nil
		}
	}
	return nil, errors.New("book not found")
}

func (m *MemoryRepository) GetAll(isbn string) (*Books, error) {
	for _, v := range m.b {
		book := Books{ISBN: v.ISBN, Name: v.Name, Status: v.Status}
		return &book, nil
	}
	return nil, errors.New("book not found")
}

// Create allows the creation of a new book into memory.
func (m *MemoryRepository) Create(isbn string, name string) (*Books, error) {
	for _, v := range m.b {
		if v.ISBN == isbn {
			return nil, errors.New("book exists")
		}
	}
	book := Books{Name: name, ISBN: isbn, Status: "Available"}
	return &book, nil
}

// Update allows the updating of book name in the memory database.
func (m *MemoryRepository) Update(isbn string, name string) (*Books, error) {
	for _, v := range m.b {
		if v.ISBN == isbn {
			book := Books{Name: name}
			return &book, nil
		}
	}
	return nil, errors.New("book cannot be updated")
}

func (m *MemoryRepository) Delete(isbn string) (*Books, error) {
	return nil, errors.New("book not found")
}
