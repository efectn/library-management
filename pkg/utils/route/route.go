package route

import (
	"github.com/efectn/library-management/pkg/middlewares/permission"
	"github.com/efectn/library-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

type ResourceController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
}

type ResourceConfig struct {
	WithPermissions bool
	Exclude         []string
}

var defaultResourceConfig = ResourceConfig{
	WithPermissions: true,
}

func checkWithPermissions(cfg ResourceConfig) func(c *fiber.Ctx) bool {
	return func(c *fiber.Ctx) bool {
		return !cfg.WithPermissions
	}
}

// The method to define route group with index, store, show, update, delete routes.
// Prefix must be singular for properly routing.
// Permissions are auto-putting. So you should create necessary permissions if you want to use permissions.
// Permission forms: list-prefix(s) (index), create-prefix (store), show-prefix(s) (show), edit-prefix (update), delete-prefix (destroy).
func CreateResource(prefix string, router fiber.Router, controller ResourceController, config ...ResourceConfig) fiber.Router {
	// Define Config
	plural := prefix + "s"
	cfg := defaultResourceConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	// Define group
	resource := router.Group("/" + plural).Name(plural + ".")

	// Define routes (index, store, show, update, delete)
	if !utils.Contains(cfg.Exclude, "index") {
		resource.Get("/", skip.New(permission.New("list-"+plural), checkWithPermissions(cfg)), controller.Index).Name("index")
	}
	if !utils.Contains(cfg.Exclude, "store") {
		resource.Post("/", skip.New(permission.New("create-"+prefix), checkWithPermissions(cfg)), controller.Store).Name("store")
	}
	if !utils.Contains(cfg.Exclude, "show") {
		resource.Get("/:id", skip.New(permission.New("show-"+plural), checkWithPermissions(cfg)), controller.Show).Name("show")
	}
	if !utils.Contains(cfg.Exclude, "update") {
		resource.Patch("/:id", skip.New(permission.New("edit-"+prefix), checkWithPermissions(cfg)), controller.Update).Name("update")
	}
	if !utils.Contains(cfg.Exclude, "delete") {
		resource.Delete("/:id", skip.New(permission.New("delete-"+prefix), checkWithPermissions(cfg)), controller.Destroy).Name("delete")
	}

	return resource
}
