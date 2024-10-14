package main

import (
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
	// Muat file .env
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// Koneksi ke database
	database.Connect()

	// Inisialisasi Fiber app
	app := fiber.New()

	// Setup routes untuk autentikasi
	handlers.AuthRoutes(app)

	// Memanggil middleware untuk token blacklist
	app.Use(middleware.TokenBlacklist)

	// Tambahkan proteksi JWT pada seluruh group branch
	branchGroup := app.Group("/api", middleware.JWTMiddleware)
	handlers.BranchRoutes(branchGroup)

	// Tambahkan proteksi JWT pada seluruh group unit
	unitGroup := app.Group("/api", middleware.JWTMiddleware)
	handlers.UnitRoutes(unitGroup)

	// Jalankan penghapusan token kadaluarsa di goroutine terpisah
	go deleteExpiredTokens()

	// Jalankan server
	app.Listen(":3000")
}
