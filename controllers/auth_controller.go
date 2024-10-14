package controllers

import (
	"fmt"
	"os"
	"strings"
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

	// Parsing request body
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	// 1. Check if user exists and status is active
	if err := database.DB.Where("username = ? AND id_branch = ?", data["username"], data["id_branch"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found or branch mismatch",
		})
	}

	// 2. Check if user is active
	if user.StatusUser != "active" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "User is inactive, please contact operator",
		})
	}

	// 3. Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		// Log failure if password is incorrect
		logFailure := models.LogFailure{
			Username:  data["username"],
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		}
		database.DB.Create(&logFailure)

		// 5. Check number of login failures in the last 24 hours
		var failureCount int64
		startOfDay := time.Now().Truncate(24 * time.Hour)
		database.DB.Model(&models.LogFailure{}).
			Where("username = ? AND timestamp >= ?", data["username"], startOfDay).
			Count(&failureCount)

		// 6. If failure count >= 3, update user status to inactive
		if failureCount >= 3 {
			user.StatusUser = "inactive"
			database.DB.Save(&user)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Too many failed login attempts, user deactivated",
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// 4. Generate JWT token if login is successful
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id_branch"] = user.IDBranch
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix() // token valid for 8 hours

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
	// Ambil token dari header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token tidak diberikan",
		})
	}

	// Hapus "Bearer " dari token jika ada
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[len("Bearer "):]
	}

	// Cek apakah token sudah ada di blacklist
	var existingToken models.TokenBlacklist
	if err := database.DB.Where("token = ?", tokenString).First(&existingToken).Error; err == nil {
		// Token sudah ada di blacklist
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Token sudah di-logout sebelumnya",
		})
	}

	// Simpan token di blacklist
	if err := database.DB.Create(&models.TokenBlacklist{
		ID:        0,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(time.Hour * 8),
	}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan token ke blacklist",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
