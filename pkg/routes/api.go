package routes

import (
	"github.com/efectn/library-management/pkg/controllers"
	"github.com/efectn/library-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var userController controllers.UserController
var authController controllers.AuthController

func RegisterAPIRoutes(app fiber.Router, config *utils.ConfigBase) {
	// Auth Routes
	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Middleware.Jwt.Key),
	}))

	// Restricted Routes
	users := app.Group("/users")

	users.Get("/", userController.Index)
	users.Post("/", userController.Store)
	users.Get("/:id", userController.Show)
	users.Patch("/:id", userController.Update)
	users.Delete("/:id", userController.Destroy)
}
