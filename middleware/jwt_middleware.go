package middleware

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/herumitra/fiberapi.git/database"
)

// JWTMiddleware is the middleware for JWT authentication
func JWTMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[len("Bearer "):]
	} else {
		log.Println("Token tidak diberikan")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token tidak diberikan"})
	}

	log.Println("Token diterima:", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println("Token tidak valid:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token tidak valid"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println("Token valid, claims:", claims)

		// Cek expiration
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().After(expirationTime) {
			log.Println("Token telah kadaluarsa")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token telah kadaluarsa"})
		}

		// Pengecekan token blacklist
		var count int64
		if err := database.DB.Table("token_blacklists").Where("token = ?", tokenString).Count(&count).Error; err != nil {
			log.Println("Error checking blacklist:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error", "catatan": database.DB.Where("token = ?", tokenString).Count(&count).Error})
		}
		if count > 0 {
			log.Println("Token telah diblacklist")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token telah diblacklist"})
		}

		c.Locals("user", claims)
		return c.Next()
	}

	log.Println("Token tidak valid")
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token tidak valid"})
}
