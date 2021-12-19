package utils

import (
	"github.com/efectn/library-management/pkg/globals"
	"gorm.io/gorm/clause"
)

func SeederFunc(model interface{}, message string, customFuncs ...func()) {
	if err := globals.App.DB.Gorm.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(model); err.Error != nil {
		globals.App.Logger.Panic().Err(err.Error).Msg("")
	}

	globals.App.DB.Gorm.Save(model)

	for _, customFunc := range customFuncs {
		customFunc()
	}

	globals.App.Logger.Info().Msg(message + " has seeded successfully!")
}
