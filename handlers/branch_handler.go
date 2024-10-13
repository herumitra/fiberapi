package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herumitra/fiberapi.git/controllers"
)

func BranchRoutes(app *fiber.App) {
	app.Post("/api/branch", controllers.CreateBranch)       // Create
	app.Get("/api/branches", controllers.GetBranches)       // Read all
	app.Get("/api/branch/:id", controllers.ShowBranch)      // Read one
	app.Put("/api/branch/:id", controllers.UpdateBranch)    // Update
	app.Delete("/api/branch/:id", controllers.DeleteBranch) // Delete
}
