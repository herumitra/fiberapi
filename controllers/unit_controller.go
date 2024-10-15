package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herumitra/fiberapi.git/helpers"
)

// Function CreateUnit
func CreateUnit(c *fiber.Ctx) error {
	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Unit created successfully",
		Data:    "", //&branch,
	}
	// Return response with status Created
	return c.Status(fiber.StatusCreated).JSON(response)
}

// Function GetUnits
func GetUnits(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

// Function ShowUnit
func ShowUnit(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

// Function UpdateUnit
func UpdateUnit(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

// Function DeleteUnit
func DeleteUnit(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
