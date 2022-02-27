package admin

import (
	"context"
	"strconv"

	"github.com/efectn/library-management/pkg/database/ent"
	erole "github.com/efectn/library-management/pkg/database/ent/role"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/utils"
	"github.com/efectn/library-management/pkg/utils/errors"
	"github.com/gofiber/fiber/v2"
)

type RoleController struct{}

type CreateRoleRequest struct {
	Name          string `validate:"required,max=32" form:"name"`
	PermissionIDs []int  `form:"permission_id" json:"permission_id,omitempty"`
}

type UpdateRoleRequest struct {
	Name          string `validate:"max=32" form:"name"`
	PermissionIDs []int  `form:"permission_id" json:"permission_id,omitempty"`
}

func (RoleController) Index(c *fiber.Ctx) error {
	roles, err := api.App.DB.Ent.Role.Query().
		WithPermissions().
		Order(ent.Asc(erole.FieldID)).
		All(context.Background())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Role list has retrieved successfully!",
		"roles":   roles,
	})
}

func (RoleController) Store(c *fiber.Ctx) error {
	r := new(CreateRoleRequest)
	if err := utils.ParseAndValidate(c, r); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	role, err := api.App.DB.Ent.Role.Create().
		SetName(r.Name).
		AddPermissionIDs(r.PermissionIDs...).
		Save(context.Background())
	if err = errors.HandleEntErrors(err); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The role created successfully!",
		"role":    role,
	})
}

func (RoleController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	role, err := api.App.DB.Ent.Role.Query().
		Where(erole.IDEQ(id)).
		WithPermissions().
		First(context.Background())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The role retrieved successfully!",
		"role":    role,
	})
}

func (RoleController) Update(c *fiber.Ctx) error {
	r := new(UpdateRoleRequest)
	if err := utils.ParseAndValidate(c, r); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	ur := api.App.DB.Ent.Role.UpdateOneID(id)

	// Update name
	if r.Name != "" {
		ur = ur.SetName(r.Name)
	}

	// Update roles
	if r.PermissionIDs != nil {
		ur = ur.ClearPermissions().AddPermissionIDs(r.PermissionIDs...)
	}

	role, err := ur.Save(context.Background())
	if err = errors.HandleEntErrors(err); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The role updated successfully!",
		"role":    role,
	})
}

func (RoleController) Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = api.App.DB.Ent.Role.DeleteOneID(id).Exec(context.Background())
	if err = errors.HandleEntErrors(err); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The role deleted successfully!",
	})
}
