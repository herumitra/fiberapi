package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herumitra/fiberapi.git/controllers"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
}
