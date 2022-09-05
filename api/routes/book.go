package routes

import (
	"github.com/amineamaach/fiber-gorm/api/handlers"
	"github.com/amineamaach/fiber-gorm/pkg/book"
	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/book/:id", handlers.GetBook(service))
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/book", handlers.AddBook(service))
	app.Put("/book/:id", handlers.UpdateBook(service))
	app.Delete("/book/:id", handlers.RemoveBook(service))
}
