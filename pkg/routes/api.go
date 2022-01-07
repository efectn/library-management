package routes

import (
	"github.com/efectn/library-management/pkg/controllers"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/middlewares/permission"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var userController controllers.UserController
var authController controllers.AuthController

func RegisterAPIRoutes(app fiber.Router) {
	// Auth Routes
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("test")
	})
	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(api.App.Config.Middleware.Jwt.Key),
	}))

	// Restricted Routes
	users := app.Group("/users")

	users.Get("/", permission.New("list-users"), userController.Index)
	users.Post("/", permission.New("create-user"), userController.Store)
	users.Get("/:id", permission.New("show-users"), userController.Show)
	users.Patch("/:id", permission.New("edit-user"), userController.Update)
	users.Delete("/:id", permission.New("delete-user"), userController.Destroy)
}
