package controllers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	database "github.com/herumitra/fiberapi.git/database"
	helpers "github.com/herumitra/fiberapi.git/helpers"
	models "github.com/herumitra/fiberapi.git/models"
	"golang.org/x/crypto/bcrypt"
)

// Generate ID user in format mitra[date][month][year][hour][minute][second]
func generateUserID() string {
	now := time.Now()
	return fmt.Sprintf("MITRA%s", now.Format("02012006150405"))
}

// Registration user
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
		BranchId:    data["branch_id"],
		StatusUser:  "active",
		Authorities: "user",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Failed to register user",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data:    "Registered user ID: " + user.ID,
	}
	// Return response with status Created
	return c.Status(fiber.StatusCreated).JSON(response)
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
	if err := database.DB.Where("username = ? AND branch_id = ?", data["username"], data["branch_id"]).First(&user).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not login user",
			Data:    "User not found or branch mismatch",
		}
		// Return response with status Not Found
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	// 2. Check if user is active
	if user.StatusUser != "active" {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not login user",
			Data:    "User is inactive, please contact operator",
		}
		// Return response with status Forbidden
		return c.Status(fiber.StatusForbidden).JSON(response)
	}

	// 3. Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		// Log failure if password is incorrect
		logFailure := models.LogFailure{
			ID:        0,
			Username:  data["username"],
			Timestamp: time.Now(),
		}
		database.DB.Create(&logFailure)

		// 4. Check number of login failures in the last 24 hours
		var failureCount int64
		startOfDay := time.Now().Truncate(24 * time.Hour)
		database.DB.Model(&models.LogFailure{}).
			Where("username = ? AND timestamp >= ?", data["username"], startOfDay).
			Count(&failureCount)

		// 5. If failure count >= 3, update user status to inactive
		if failureCount >= 3 {
			user.StatusUser = "inactive"
			database.DB.Save(&user)
			// Set format response
			response := helpers.Response{
				Status:  "failure",
				Message: "Could not login user",
				Data:    "Too many failed login attempts, user deactivated",
			}
			// Return response with status Forbidden
			return c.Status(fiber.StatusForbidden).JSON(response)
		}

		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not login user",
			Data:    "Error: Invalid credentials",
		}
		// Return response with status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// 6. Generate JWT token if login is successful
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["branch_id"] = user.BranchId
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix() // token valid for 8 hours

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not login",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Login successfully",
		Data:    "Bearer " + t,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Logout user
func Logout(c *fiber.Ctx) error {
	// Ambil token dari header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 	"message": "Token tidak diberikan",
		// })
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not logout user",
			Data:    "Error: Token not provided",
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// Hapus "Bearer " dari token jika ada
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[len("Bearer "):]
	}

	// Cek apakah token sudah ada di blacklist
	var existingToken models.TokenBlacklist
	if err := database.DB.Where("token = ?", tokenString).First(&existingToken).Error; err == nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could logout user",
			Data:    "Error: Token already logged out",
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusConflict).JSON(response)
	}

	// Simpan token di blacklist
	if err := database.DB.Create(&models.TokenBlacklist{
		ID:        0,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(time.Hour * 8),
	}).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not logout user",
			Data:    "Error: Failed to save token to blacklist",
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Logged out successfully",
		Data:    "User logged out successfully",
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}
