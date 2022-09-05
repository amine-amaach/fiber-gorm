package presenter

import (
	"github.com/amineamaach/fiber-gorm/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

// Book is the presenter object which will be passed in the response by Handler
type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func BookSuccessResponse(data *entities.Book) *fiber.Map {
	book := Book{
		ID:     data.ID,
		Title:  data.Title,
		Author: data.Author,
		Rating: data.Rating,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

// BooksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func BooksSuccessResponse(data *[]entities.Book) *fiber.Map {
	var books []Book
	for _, v := range *data {
		books = append(books, Book{
			ID:     v.ID,
			Title:  v.Title,
			Author: v.Author,
			Rating: v.Rating,
		})
	}
	return &fiber.Map{
		"status": true,
		"data":   books,
		"error":  nil,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
