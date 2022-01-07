package seeds

import (
	"context"
	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/efectn/library-management/pkg/globals/api"
)

type UserSeeder struct{}

var users = []ent.User{
	{
		Email:    "john@test.net",
		Password: "12345",
		Name:     "John Doe",
		Phone:    "+905000000000",
		ZipCode:  11111,
	},
	{
		Email:    "jane@test.net",
		Password: "12345",
		Name:     "Jane Doe",
		Phone:    "+902000000000",
		ZipCode:  22222,
	},
}

func (UserSeeder) Seed() error {
	bulk := make([]*ent.UserCreate, len(users))
	for i, user := range users {
		bulk[i] = api.App.DB.Ent.User.Create().SetEmail(user.Email).
			SetPassword(user.Password).
			SetName(user.Name).
			SetPhone(user.Phone).
			SetZipCode(user.ZipCode)
	}
	_, err := api.App.DB.Ent.User.CreateBulk(bulk...).Save(context.Background())
	if err != nil {
		return err
	}

	return nil

	/*utils.SeederFunc(&users, "Users", func() {
		utils.Authority{}.AssignRole(users[0].ID, "Admin")
		utils.Authority{}.AssignRole(users[1].ID, "User")
	})*/
}

func (UserSeeder) Count() (int, error) {
	return api.App.DB.Ent.User.Query().Count(context.Background())
}
