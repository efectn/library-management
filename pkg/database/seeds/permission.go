package seeds

import (
	"context"
	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/efectn/library-management/pkg/globals/api"
)

type PermissionSeeder struct{}

var permissions = []ent.Role{
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

func (PermissionSeeder) Seed() error {
	bulk := make([]*ent.PermissionCreate, len(permissions))
	for i, perm := range permissions {
		bulk[i] = api.App.DB.Ent.Permission.Create().SetName(perm.Name)
	}
	_, err := api.App.DB.Ent.Permission.CreateBulk(bulk...).Save(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (PermissionSeeder) Count() (int, error) {
	return api.App.DB.Ent.Permission.Query().Count(context.Background())
}
