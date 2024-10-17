package middleware

import (
	validator "github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	helpers "github.com/herumitra/fiberapi.git/helpers"
	models "github.com/herumitra/fiberapi.git/models"
)

// ValidateUnit is a function to validate struct Data of Unit
func ValidateUnit(unit models.Unit) []*helpers.ErrorResponse {
	var validate = validator.New()
	var errors []*helpers.ErrorResponse
	err := validate.Struct(unit)

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

// ValidateUnitField is a middleware to validate field of Unit
func ValidateUnitField(c *fiber.Ctx) error {

	unit := new(models.Unit)

	if err := c.BodyParser(unit); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	errors := ValidateUnit(*unit)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	//Return Next Function
	return c.Next()
}
