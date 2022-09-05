# Clean Architecture with GoFiber and Gorm/postgres

A simple book shop CRUD demonstration implementing the Clean Architecture in GoFiber with Gorm/postgres

## Installation

In order to run the project, please follow the following steps:

1. Clone the Repo
2. Go to the "gorm" folder
3. Run `go get`
4. Replace your Postgres DB Connection string in `app.go`

## Routes

|     API Path   | Method |               What it does              |
|:--------------:|:------:|:---------------------------------------:|
| /api/books     |   GET  | Fetches the list of books from the shop |
| /api/book/{id} |   GET  |        Fetches book from the shop       |
| /api/book      |  POST  |      Creates/Adds book to the shop      |
| /api/book/{id} | DELETE |    Removes/Deletes book from the shop   |
| /api/boos/{id} |  PUT   |  Updates the book details from the shop |
