package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/herumitra/fiberapi.git/database"
)

func TokenBlacklist(c *fiber.Ctx) error {
	// Ambil token dari header
	tokenString := c.Get("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[len("Bearer "):]
	} else {
		log.Println("Token tidak diberikan")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token tidak diberikan"})
	}

	// Cek apakah token ada di blacklist
	var tokenBlacklisted bool
	if err := database.DB.Table("token_blacklists").Where("token = ?", tokenString).First(&tokenBlacklisted).Error; err == nil && tokenBlacklisted {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token telah dibatalkan",
		})
	}

	// Lanjutkan ke middleware berikutnya
	return c.Next()
}
