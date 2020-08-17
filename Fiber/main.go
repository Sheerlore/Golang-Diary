package main

import (
	"fmt"

	"github.com/Sheerlore/GolngMytools/Fiber/book"
	"github.com/Sheerlore/GolngMytools/Fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


func setupRoutes(app *fiber.App) {

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBconn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}



func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	app.Listen(3000)
	defer database.DBconn.Close()
}