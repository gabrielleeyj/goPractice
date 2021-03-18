package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"library.com/library/models"
)

// Repository interface is the implementation of the Repository functions.
type Repository interface {
	Get(isbn string) (*models.Book, error)
	GetAll() ([]models.Book, error)
	Create(book models.Book) (*models.Book, error)
	Update(id uint, book models.Book) (*models.Book, error)
	Delete(id uint) error
}

// BookRepository is the implementation of the database via Gorm
type BookRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &BookRepository{
		db: db,
	}
}
func (r *BookRepository) Get(isbn string) (*models.Book, error) {
	var book models.Book
	err := r.db.Model(&models.Book{}).Where("isbn = ?", isbn).First(&book).Error //First not Find
	if err != nil {
		return nil, errors.New("book not found")
	}
	return &book, nil
}

// GetAll
func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Model(&models.Book{}).Find(&books).Error
	if err != nil {
		return nil, errors.New("no books found")
	}
	return books, nil
}

func (r *BookRepository) Create(book models.Book) (*models.Book, error) {
	err := r.db.Model(&models.Book{}).Create(&book).Error
	if err != nil {
		return nil, errors.New("book exists")
	}
	return &book, nil
}

// isbn, data structure of entity
func (r *BookRepository) Update(id uint, book models.Book) (*models.Book, error) {
	query := r.db.Model(&models.Book{}).Where("id = ?", id).Updates(&book)
	if query.Error != nil || query.RowsAffected >= 1 {
		return nil, errors.New("book cannot be updated")
	}
	return &book, nil
}

func (r *BookRepository) Delete(id uint) error {
	err := r.db.Where("id = ?", id).Delete(&models.Book{}).Error
	if err != nil {
		return errors.New("book not found")
	}
	return nil
}

/*													*/
/* 													*/
/*			MEMORY REPOSITORY CODE BLOCK			*/
/*													*/
/*													*/

// MemoryRepository is the implementation of data storage in memory.
type MemoryRepository struct {
	b []models.Book
}

// NewMemoryRepository stores an array of books
func NewMemoryRepository(b []models.Book) Repository {
	return &MemoryRepository{
		b: b,
	}
}

// Get searches for a single book in memory
func (m *MemoryRepository) Get(isbn string) (*models.Book, error) {
	for _, v := range m.b {
		if v.ISBN == isbn {
			return &models.Book{Name: v.Name}, nil
		}
	}
	return nil, errors.New("book not found")
}

// GetAll gets all the books in memory
func (m *MemoryRepository) GetAll() ([]models.Book, error) {
	if m.b == nil {
		return nil, errors.New("book not found")
	}
	for _, v := range m.b {
		fmt.Println("Name: ", v.Name)
		fmt.Println("ISBN: ", v.ISBN)
		fmt.Println("Status: ", v.Status)
	}
	return m.b, nil
}

// Create allows the creation of a new book into memory.
func (m *MemoryRepository) Create(book models.Book) (*models.Book, error) {
	// checking function
	for _, v := range m.b {
		if v.ISBN == book.ISBN {
			return nil, errors.New("book exists")
		}
	}
	//if ValidateISBN(book.ISBN) != false {
	m.b = append(m.b, book)
	//}
	return &book, nil
}

// Update allows the updating of book name in the memory database.
func (m *MemoryRepository) Update(id uint, book models.Book) (*models.Book, error) {
	for _, v := range m.b {
		if v.ISBN == book.ISBN {
			book := models.Book{Name: book.Name}
			return &book, nil
		}
	}
	return nil, errors.New("book cannot be updated")
}

// Delete checks whether isbn exist and deletes
func (m *MemoryRepository) Delete(id uint) error {
	temp := make([]models.Book, 0)
	for _, v := range m.b {
		if v.ID != id {
			temp = append(temp, v)
		}
	}

	return errors.New("book does not exist")
}
