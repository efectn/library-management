package seeds

import (
	"context"

	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/utils"
)

type RoleSeeder struct{}

var roles = []ent.Role{
	{
		Name: "User",
	},
	{
		Name: "Admin",
	},
}

func (RoleSeeder) Seed() error {
	bulk := make([]*ent.RoleCreate, len(roles))
	for i, role := range roles {
		bulk[i] = api.App.DB.Ent.Role.Create().SetName(role.Name)
	}
	roles, err := api.App.DB.Ent.Role.CreateBulk(bulk...).Save(context.Background())
	if err != nil {
		return err
	}

	// Assign permissions
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

	return nil
}

func (RoleSeeder) Count() (int, error) {
	return api.App.DB.Ent.Role.Query().Count(context.Background())
}
