package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	database "github.com/herumitra/fiberapi.git/database"
	helpers "github.com/herumitra/fiberapi.git/helpers"
	models "github.com/herumitra/fiberapi.git/models"
)

// Function generateUnitID
func generateUnitID() string {
	now := time.Now()
	return fmt.Sprintf("UNT%s", now.Format("02012006150405"))
}

// Function CreateUnit
func CreateUnit(c *fiber.Ctx) error {
	var unit models.Unit

	if err := c.BodyParser(&unit); err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Cannot parse JSON",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	unit.ID = generateUnitID()

	if err := database.DB.Create(&unit).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not create unit",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Unit created successfully",
		Data:    &unit,
	}
	// Return response with status Created
	return c.Status(fiber.StatusCreated).JSON(response)
}

// Function GetUnits
func GetUnits(c *fiber.Ctx) error {
	var units []models.Unit
	if err := database.DB.Find(&units).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve units",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Units retrieved successfully",
		Data:    &units,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Function ShowUnit
func ShowUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var unit models.Unit
	if err := database.DB.First(&unit, "id = ?", id).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve units",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Unit retrieved successfully",
		Data:    &unit,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Function UpdateUnit
func UpdateUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var unit models.Unit
	if err := database.DB.First(&unit, "id = ?", id).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve units",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	if err := c.BodyParser(&unit); err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Cannot parse JSON",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	if err := database.DB.Save(&unit).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not update unit",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Unit updated successfully",
		Data:    &unit,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Function DeleteUnit
func DeleteUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var unit models.Unit
	if err := database.DB.First(&unit, "id = ?", id).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve units",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	if err := database.DB.Delete(&unit).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not delete unit",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Unit deleted successfully",
		Data:    &unit,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}
