package middleware

import (
	validator "github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	helpers "github.com/herumitra/fiberapi.git/helpers"
	models "github.com/herumitra/fiberapi.git/models"
)

// ValidateSupplier is a function to validate struct Data of Supplier
func ValidateSupplier(supplier models.Supplier) []*helpers.ErrorResponse {
	var validate = validator.New()
	var errors []*helpers.ErrorResponse
	err := validate.Struct(supplier)

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

// ValidateSupplierField is a middleware to validate field of Supplier
func ValidateSupplierField(c *fiber.Ctx) error {

	supplier := new(models.Supplier)

	if err := c.BodyParser(supplier); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	errors := ValidateSupplier(*supplier)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	//Return Next Function
	return c.Next()
}
