package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/herumitra/fiberapi.git/database"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	branch.ID = generateBranchID()

	if err := database.DB.Create(&branch).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create branch",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Branch created successfully",
		"branch":  branch,
	})
}

// Get all Branches
func GetBranches(c *fiber.Ctx) error {
	var branches []models.Branch

	if err := database.DB.Find(&branches).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not retrieve branches",
			"error":   err.Error(),
		})
	}

	return c.JSON(branches)
}

// Get single Branch by ID
func ShowBranch(c *fiber.Ctx) error {
	id := c.Params("id")
	var branch models.Branch

	if err := database.DB.First(&branch, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Branch not found",
		})
	}

	return c.JSON(branch)
}

// Update Branch
func UpdateBranch(c *fiber.Ctx) error {
	id := c.Params("id")
	var branch models.Branch

	if err := database.DB.First(&branch, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Branch not found",
		})
	}

	if err := c.BodyParser(&branch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	if err := database.DB.Save(&branch).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update branch",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Branch updated successfully",
		"branch":  branch,
	})
}

// Delete Branch
func DeleteBranch(c *fiber.Ctx) error {
	id := c.Params("id")
	var branch models.Branch

	if err := database.DB.First(&branch, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Branch not found",
		})
	}

	if err := database.DB.Delete(&branch).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete branch",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Branch deleted successfully",
	})
}
