package utils

import (
	"fmt"

	"github.com/efectn/library-management/pkg/app"
	"gorm.io/gorm/clause"
)

func SeederFunc(model interface{}, message string, customFuncs ...func()) {
	if err := app.App.DB.Gorm.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(model); err.Error != nil {
		fmt.Printf("\n=====> ERROR: %v\n", err.Error)
	}

	app.App.DB.Gorm.Save(model)

	for _, customFunc := range customFuncs {
		customFunc()
	}

	fmt.Println("=====> INFO: " + message + " has seeded successfully!")
}
