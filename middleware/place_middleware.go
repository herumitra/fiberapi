package middleware

import (
	validator "github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	helpers "github.com/herumitra/fiberapi.git/helpers"
	models "github.com/herumitra/fiberapi.git/models"
)

func ValidatePlace(place models.Place) []*helpers.ErrorResponse {
	var validate = validator.New()
	var errors []*helpers.ErrorResponse
	err := validate.Struct(place)

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

func ValidatePlaceField(c *fiber.Ctx) error {

	place := new(models.Place)

	if err := c.BodyParser(place); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	errors := ValidatePlace(*place)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	//Return Next Function
	return c.Next()
}
