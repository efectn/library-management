package routes

import (
	"github.com/efectn/library-management/pkg/controllers"
	"github.com/efectn/library-management/pkg/controllers/admin"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/utils/route"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var userController admin.UserController
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
		SigningKey:   []byte(api.App.Config.Middleware.Jwt.Key),
		ErrorHandler: api.App.Fiber.ErrorHandler,
	}))

	// Restricted Routes
	// Admin Routes

	admin := app.Group("/admin").Name("admin.")
	route.CreateResource("user", admin, userController)
}
