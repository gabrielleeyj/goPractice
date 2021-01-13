package main

import "gorm.io/gorm"

// Repository interface is the implementation of the Repository functions.
type Repository interface {
	Get(isbn string) (*Book, error)
	GetAll(isbn string) (*Book, error)
	Create(isbn, name string) (*Book, error)
	Update(isbn, name string) (*Book, error)
	Delete(isbn string) error
}

// BookRespository is the implementation of the database
type BookRespository struct {
	db *gorm.DB
}

func (r *BookRespository) Get(isbn string) (*Book, error) {

}

func (r *BookRespository) GetAll(isbn string) (*Book, error) {

}

func (r *BookRespository) Create(isbn string, name string) (*Book, error) {

}

func (r *BookRespository) Update(isbn string, name string) (*Book, error) {

}

func (r *BookRespository) Delete(isbn string) (*Book, error) {

}
