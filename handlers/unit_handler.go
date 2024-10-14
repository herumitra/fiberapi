package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herumitra/fiberapi.git/controllers"
)

func UnitRoutes(group fiber.Router) {
	group.Get("/units", controllers.GetUnits)
	group.Get("/unit/:id", controllers.ShowUnit)
	group.Post("/unit", controllers.CreateUnit)
	group.Put("/unit/:id", controllers.UpdateUnit)
	group.Delete("/unit/:id", controllers.DeleteUnit)
}
