package seeds

import (
	"github.com/efectn/library-management/pkg/database/models"
	"github.com/efectn/library-management/pkg/utils"
)

type PermissionSeeder struct{}

var permissions = []models.Permission{
	{
		Name: "access-profile",
	},
	{
		Name: "list-users",
	},
	{
		Name: "show-users",
	},
	{
		Name: "create-user",
	},
	{
		Name: "edit-user",
	},
	{
		Name: "delete-user",
	},
	{
		Name: "list-roles",
	},
	{
		Name: "show-roles",
	},
	{
		Name: "create-role",
	},
	{
		Name: "edit-role",
	},
	{
		Name: "delete-role",
	},
}

func (PermissionSeeder) Seed() {
	utils.SeederFunc(&permissions, "Permissions")
}

func (PermissionSeeder) ReturnModel() interface{} {
	return &models.Permission{}
}
