package bookcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herujci/fiberapi.git/helpers"
	"github.com/herujci/fiberapi.git/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	// Declare variable books of type Book from models
	var books []models.Book

	// Condition when error with status internal server error
	if err := models.DB.Find(&books).Error; err != nil {
		// Create response failure created using format from helpers
		response := helpers.Response{
			Status:  "failure",
			Message: "Books not found",
			Data:    "Error: " + err.Error(),
		}
		// Return response Status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Get all books using format from helpers
	response := helpers.Response{
		Status:  "success",
		Message: "Books retrieved successfully",
		Data:    books,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)
}

func Show(c *fiber.Ctx) error {

	// Get id from request parameter
	id := c.Params("id")

	// Declare variable book of type Book from models
	var book models.Book

	// Condition when error with status not found or Internal Server Error
	if err := models.DB.First(&book, id).Error; err != nil {
		// Condition when error with status not found
		if err == gorm.ErrRecordNotFound {
			// Create response failure created using format from helpers
			response := helpers.Response{
				Status:  "failure",
				Message: "Book not found",
				Data:    "Error: " + err.Error(),
			}
			// Return response Status Not Found
			return c.Status(fiber.StatusNotFound).JSON(response)
		}

		// Condition when error with status internal server error
		if err == gorm.ErrInvalidData {
			// Create response failure created using format from helpers
			response := helpers.Response{
				Status:  "failure",
				Message: "Book not found",
				Data:    "Error: " + err.Error(),
			}
			// Return response Status Internal Server Error
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}

	// Get book by id
	response := helpers.Response{
		Status:  "success",
		Message: "Books retrieved successfully",
		Data:    book,
	}
	// Return response with status OK
	return c.Status(fiber.StatusOK).JSON(response)

}

func Create(c *fiber.Ctx) error {

	// Declare variable book of type Book from models
	var book models.Book

	// Condition when error with status bad request
	if err := c.BodyParser(&book); err != nil {
		// Create response failure created using format from helpers
		response := helpers.Response{
			Status:  "failure",
			Message: "Book insert failed",
			Data:    "Error: " + err.Error(),
		}
		// Return response Status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Condition when error with status internal server error
	if err := models.DB.Create(&book).Error; err != nil {
		// Create response failure created using format from helpers
		response := helpers.Response{
			Status:  "failure",
			Message: "Book insert failed",
			Data:    "Error: " + err.Error(),
		}
		// Return response Status Internal Server Error
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Get book by id and response success created using format from helpers
	response := helpers.Response{
		Status:  "success",
		Message: "Book inserted successfully",
		Data:    book,
	}
	// Return response Status Created
	return c.Status(fiber.StatusCreated).JSON(response)
}

func Update(c *fiber.Ctx) error {
	// Get id from request parameter
	id := c.Params("id")

	// Declare variable book of type Book from models
	var book models.Book

	// Condition when error with status bad request
	if err := c.BodyParser(&book); err != nil {
		// Create response failure created using format from helpers
		response := helpers.Response{
			Status:  "failure",
			Message: "Book update failed",
			Data:    "Error: " + err.Error(),
		}
		// Return response Status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Return response with status bad request
	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		// Create response failure created using format from helpers
		response := helpers.Response{
			Status:  "failure",
			Message: "Book update failed",
			Data:    "Error: Can not update book",
		}
		// Return response Status Internal Server Error
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Update book by id
	response := helpers.Response{
		Status:  "success",
		Message: "Book updated successfully",
		Data:    book,
	}

	// Return response with status ok
	return c.Status(fiber.StatusOK).JSON(response)
}

func Delete(c *fiber.Ctx) error {
	// Get id from request parameter
	id := c.Params("id")

	// Declare variable book of type Book from models
	var book models.Book

	if models.DB.Delete(&book, id).RowsAffected == 0 {
		// Create response failure created using format from helpers
		response := helpers.Response{
			Status:  "failure",
			Message: "Book delete failed",
			Data:    "Error: Can not delete book",
		}
		// Return response Status Not Found
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	// Delete book by id
	response := helpers.Response{
		Status:  "success",
		Message: "Book deleted successfully",
		Data:    "Deleted book successfully",
	}
	// Return response with status ok
	return c.Status(fiber.StatusOK).JSON(response)

}
