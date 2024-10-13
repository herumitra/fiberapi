package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herumitra/fiberapi.git/database"
	"github.com/herumitra/fiberapi.git/handlers"
)

func main() {
	// Koneksi ke database
	database.Connect()

	// Inisialisasi Fiber app
	app := fiber.New()

	// Setup routes
	handlers.AuthRoutes(app)
	handlers.BranchRoutes(app)

	// Jalankan server
	app.Listen(":3000")
}
