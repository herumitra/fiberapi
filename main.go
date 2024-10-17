package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/herumitra/fiberapi.git/database"
	"github.com/herumitra/fiberapi.git/handlers"
	"github.com/herumitra/fiberapi.git/middleware"
	"github.com/herumitra/fiberapi.git/models"
	"github.com/joho/godotenv"
)

// Fungsi untuk menghapus token kadaluarsa
func deleteExpiredTokens() {
	for {
		time.Sleep(1 * time.Hour) // Menghapus setiap jam
		database.DB.Where("expires_at < ?", time.Now()).Delete(&models.TokenBlacklist{})
	}
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// Inialisasi database
	database.Connect()

	// Inisialisasi Fiber app
	app := fiber.New()

	// Setup routes auth
	handlers.AuthRoutes(app)

	// Setup middleware token blacklist
	app.Use(middleware.TokenBlacklist)

	// Add protection JWT for group branch
	branchGroup := app.Group("/api", middleware.JWTMiddleware)
	handlers.BranchRoutes(branchGroup)

	// Add protection JWT for group unit
	unitGroup := app.Group("/api", middleware.JWTMiddleware)
	handlers.UnitRoutes(unitGroup)

	// Running delete expired tokens in separate goroutine
	go deleteExpiredTokens()

	// Jalankan server
	app.Listen(":" + os.Getenv("SERVER_PORT"))
}
