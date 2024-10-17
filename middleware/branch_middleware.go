package middleware

import (
	validator "github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	helpers "github.com/herumitra/fiberapi.git/helpers"
	models "github.com/herumitra/fiberapi.git/models"
)

// ValidateBranch is a function to validate struct Data of Branch
func ValidateBranch(branch models.Branch) []*helpers.ErrorResponse {
	var validate = validator.New()
	var errors []*helpers.ErrorResponse
	err := validate.Struct(branch)

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

// ValidateBranchField is a middleware to validate field of Branch
func ValidateBranchField(c *fiber.Ctx) error {

	branch := new(models.Branch)

	if err := c.BodyParser(branch); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	errors := ValidateBranch(*branch)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	//Return Next Function
	return c.Next()
}
