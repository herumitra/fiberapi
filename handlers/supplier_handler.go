package handlers

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/herumitra/fiberapi.git/controllers"
)

func SupplierRoutes(app *fiber.App) {
	app.Get("/suppliers", controllers.GetSuppliers)
	app.Get("/supplier/:id", controllers.ShowSupplier)
	app.Post("/supplier", controllers.CreateSupplier)
	app.Put("/supplier/:id", controllers.UpdateSupplier)
	app.Delete("/supplier/:id", controllers.DeleteSupplier)
}
