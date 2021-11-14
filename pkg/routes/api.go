package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ofsahof/library-management/pkg/controllers"
)

var userController controllers.UserController

func RegisterAPIRoutes(app fiber.Router) {
	users := app.Group("/users")

	users.Get("/", userController.Index)
	users.Post("/", userController.Store)
	users.Get("/:id", userController.Show)
	users.Patch("/:id", userController.Update)
	users.Delete("/:id", userController.Destroy)
}
