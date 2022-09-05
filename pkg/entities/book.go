package entities

import "gorm.io/gorm"

// Book Constructs your Book model under entities.
type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}