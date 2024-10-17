package handlers

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/herumitra/fiberapi.git/controllers"
	middleware "github.com/herumitra/fiberapi.git/middleware"
)

func UnitRoutes(group fiber.Router) {
	group.Get("/units", middleware.ValidateUnitField, controllers.GetUnits)
	group.Get("/unit/:id", middleware.ValidateUnitField, controllers.ShowUnit)
	group.Post("/unit", middleware.ValidateUnitField, controllers.CreateUnit)
	group.Put("/unit/:id", middleware.ValidateUnitField, controllers.UpdateUnit)
	group.Delete("/unit/:id", middleware.ValidateUnitField, controllers.DeleteUnit)
}
