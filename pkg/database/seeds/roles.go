package seeds

import (
	"context"
	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/efectn/library-management/pkg/globals/api"
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
	_, err := api.App.DB.Ent.Role.CreateBulk(bulk...).Save(context.Background())
	if err != nil {
		return err
	}

	return nil

	/*utils.SeederFunc(&roles, "Roles", func() {
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
	})*/
}

func (RoleSeeder) Count() (int, error) {
	return api.App.DB.Ent.Role.Query().Count(context.Background())
}
