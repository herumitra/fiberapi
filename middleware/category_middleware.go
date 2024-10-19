package middleware

import (
	validator "github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	helpers "github.com/herumitra/fiberapi.git/helpers"
	models "github.com/herumitra/fiberapi.git/models"
)

// ValidateCategory is a function to validate struct Data of Category
func ValidateCategory(category models.Category) []*helpers.ErrorResponse {
	var validate = validator.New()
	var errors []*helpers.ErrorResponse
	err := validate.Struct(category)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

// ValidateCategoryField is a middleware to validate field of Category
func ValidateCategoryField(c *fiber.Ctx) error {

	category := new(models.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	errors := ValidateCategory(*category)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	//Return Next Function
	return c.Next()
}
