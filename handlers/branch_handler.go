package handlers

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/herumitra/fiberapi.git/controllers"
)

func BranchRoutes(group fiber.Router) {
	// Semua route branch akan dilindungi oleh JWT
	group.Get("/branches", controllers.GetBranches)       // Read all branches
	group.Get("/branch/:id", controllers.ShowBranch)      // Read one branch
	group.Post("/branch", controllers.CreateBranch)       // Create a new branch
	group.Put("/branch/:id", controllers.UpdateBranch)    // Update a branch
	group.Delete("/branch/:id", controllers.DeleteBranch) // Delete a branch
}
