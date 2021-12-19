package seeds

import (
	"github.com/efectn/library-management/pkg/database/models"
	"github.com/efectn/library-management/pkg/utils"
)

type RoleSeeder struct{}

var roles = []models.Role{
	{
		Name: "User",
	},
	{
		Name: "Admin",
	},
}

func (RoleSeeder) Seed() {
	utils.SeederFunc(&roles, "Roles", func() {
		utils.Authority{}.AssignPermissions(roles[0].Name, "access-profile")
		utils.Authority{}.AssignPermissions(roles[1].Name,
			"access-profile",
			"list-users",
			"show-users",
			"create-user",
			"edit-user",
			"delete-user",
			"list-roles",
			"show-roles",
			"create-role",
			"edit-role",
			"delete-role",
		)
	})
}

func (RoleSeeder) ReturnModel() interface{} {
	return &models.Role{}
}
