package main

import (
	"fmt"
	"log"

	"github.com/amineamaach/fiber-gorm/api/routes"
	"github.com/amineamaach/fiber-gorm/pkg/book"
	"github.com/amineamaach/fiber-gorm/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "password"
	dbname   = "demo"
)

func main() {
	db, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error : ", err)
	}
	fmt.Println("Database connection success!")
	db.AutoMigrate(&entities.Book{})

	bookRepo := book.NewRepo(db)
	bookSvc := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture postgres book shop!"))
	})
	api := app.Group("/api/v1")
	routes.BookRouter(api, bookSvc)
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
