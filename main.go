package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herujci/fiberapi.git/controllers/bookcontroller"
	"github.com/herujci/fiberapi.git/models"
)

func main() {
	models.ConnectDB()
	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("/", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

	app.Listen(":1199")
}
