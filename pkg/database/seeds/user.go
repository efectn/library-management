package seeds

import (
	"github.com/efectn/library-management/pkg/database/models"
	"github.com/efectn/library-management/pkg/utils"
)

type UserSeeder struct{}

var users = []models.Users{
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

func (UserSeeder) Seed() {
	utils.SeederFunc(&users, "Users")
}

func (UserSeeder) ReturnModel() interface{} {
	return &models.Users{}
}
