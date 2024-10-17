package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	database "github.com/herumitra/fiberapi.git/database"
	helpers "github.com/herumitra/fiberapi.git/helpers"
	models "github.com/herumitra/fiberapi.git/models"
)

// Function generateSupplierID
func GenerateSupplierID() string {
	now := time.Now()
	return fmt.Sprintf("SPL%s", now.Format("02012006150405"))
}

// Create Supplier
func CreateSupplier(c *fiber.Ctx) error {
	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Cannot parse JSON",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	supplier.ID = GenerateSupplierID()
	if err := database.DB.Create(&supplier).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not create supplier",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Supplier created successfully",
		Data:    &supplier,
	}
	// Return response with status Created
	return c.Status(fiber.StatusCreated).JSON(response)
}

// Get Suppliers
func GetSuppliers(c *fiber.Ctx) error {
	var suppliers []models.Supplier
	if err := database.DB.Find(&suppliers).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve suppliers",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Suppliers retrieved successfully",
		Data:    &suppliers,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Get single Supplier by ID
func ShowSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier
	if err := database.DB.First(&supplier, "id = ?", id).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve suppliers",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Supplier retrieved successfully",
		Data:    &supplier,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Update Supplier
func UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier
	if err := database.DB.First(&supplier, "id = ?", id).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve suppliers",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	if err := c.BodyParser(&supplier); err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Cannot parse JSON",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	if err := database.DB.Save(&supplier).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not update supplier",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Supplier updated successfully",
		Data:    &supplier,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Delete Supplier
func DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier
	if err := database.DB.First(&supplier, "id = ?", id).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve suppliers",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	if err := database.DB.Delete(&supplier).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not delete supplier",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Supplier deleted successfully",
		Data:    &supplier,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}
