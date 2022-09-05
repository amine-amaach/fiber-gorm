package book

import (
	"github.com/amineamaach/fiber-gorm/pkg/entities"
	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations in pg.
type Repository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	ReadBook(ID uint) (*entities.Book, error)
	ReadBooks() (*[]entities.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(ID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {
	return book, r.db.Create(book).Error
}

func (r *repository) ReadBook(ID uint) (*entities.Book, error) {
	var book entities.Book
	result := r.db.Find(&book, ID)
	return &book, result.Error
}

func (r *repository) ReadBooks() (*[]entities.Book, error) {
	var books []entities.Book
	result := r.db.Find(&books)
	return &books, result.Error
}

func (r *repository) UpdateBook(book *entities.Book) (*entities.Book, error) {
	return book, r.db.Save(book).Error
}

func (r *repository) DeleteBook(ID uint) error {
	var book entities.Book
	if err := r.db.Find(&book, ID).Error; err != nil {
		return err
	}
	return r.db.Delete(&book).Error
}
