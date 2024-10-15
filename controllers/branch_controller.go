package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/herumitra/fiberapi.git/database"
	"github.com/herumitra/fiberapi.git/helpers"
	"github.com/herumitra/fiberapi.git/models"
)

// Fungsi untuk generate ID branch dengan format branch[tanggal][bulan][tahun][jam][menit][detik]
func generateBranchID() string {
	now := time.Now()
	return fmt.Sprintf("BR%s", now.Format("02012006150405"))
}

// Create Branch
func CreateBranch(c *fiber.Ctx) error {
	var branch models.Branch

	if err := c.BodyParser(&branch); err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Cannot parse JSON",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	branch.ID = generateBranchID()

	if err := database.DB.Create(&branch).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not create branch",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Branch created successfully",
		Data:    &branch,
	}
	// Return response with status Created
	return c.Status(fiber.StatusCreated).JSON(response)
}

// Get all Branches
func GetBranches(c *fiber.Ctx) error {
	var branches []models.Branch

	if err := database.DB.Find(&branches).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve branches",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Branch retrieved successfully",
		Data:    &branches,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Get single Branch by ID
func ShowBranch(c *fiber.Ctx) error {
	id := c.Params("id")
	var branch models.Branch

	if err := database.DB.First(&branch, "id = ?", id).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve branches",
			Data:    "Branch not found",
		}
		// Return response with status Not Found
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Branch retrieved successfully",
		Data:    &branch,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Update Branch
func UpdateBranch(c *fiber.Ctx) error {
	id := c.Params("id")
	var branch models.Branch

	if err := database.DB.First(&branch, "id = ?", id).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve branches",
			Data:    "Branch not found",
		}
		// Return response with status Not Found
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	if err := c.BodyParser(&branch); err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Cannot parse JSON",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := database.DB.Save(&branch).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not update branch",
			Data:    "Error: " + err.Error(),
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Branch update successfully",
		Data:    &branch,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

// Delete Branch
func DeleteBranch(c *fiber.Ctx) error {
	id := c.Params("id")
	var branch models.Branch

	if err := database.DB.First(&branch, "id = ?", id).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve branches",
			Data:    "Branch not found",
		}
		// Return response with status Not Found
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	if err := database.DB.Delete(&branch).Error; err != nil {
		// Set format response
		response := helpers.Response{
			Status:  "failure",
			Message: "Could not retrieve branches",
			Data:    "Branch not found",
		}
		// Return response with status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Set format response
	response := helpers.Response{
		Status:  "success",
		Message: "Branch deleted successfully",
		Data:    "Deleted id: " + id,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}
