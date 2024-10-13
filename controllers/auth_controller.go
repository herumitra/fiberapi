package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/herumitra/fiberapi.git/database"
	"github.com/herumitra/fiberapi.git/models"
	"golang.org/x/crypto/bcrypt"
)

// Fungsi untuk generate ID user dengan format mitra[tanggal][bulan][tahun][jam][menit][detik]
func generateUserID() string {
	now := time.Now()
	return fmt.Sprintf("MITRA%s", now.Format("02012006150405"))
}

// Register user baru
func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		ID:          generateUserID(),
		Username:    data["username"],
		Password:    string(hashedPassword),
		IDBranch:    data["id_branch"],
		StatusUser:  "active",
		Authorities: "user",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}

// Login user
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	if err := database.DB.Where("username = ? AND id_branch = ?", data["username"], data["id_branch"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		logFailure := models.LogFailure{
			Username:  data["username"],
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		}
		database.DB.Create(&logFailure)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id_branch"] = user.IDBranch
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"token": t,
	})
}

// Logout user
func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
