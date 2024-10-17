package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GetBranchIDFromToken extracts branch_id from the JWT token
func GetBranchIDFromToken(c *fiber.Ctx) (string, error) {
	// Get the token from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Authorization header missing")
	}

	// Split the header to get the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid authorization format")
	}

	// Parse the token
	tokenString := parts[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Extract branch_id from claims
	idBranch, ok := claims["branch_id"].(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "branch_id not found in token")
	}

	return idBranch, nil
}

// Function GetBranchId
func GetBranchId(c *fiber.Ctx) string {
	BranchId, _ := GetBranchIDFromToken(c)
	return fmt.Sprintf("%s", BranchId)
}
