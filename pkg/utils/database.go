package utils

import (
	"github.com/efectn/library-management/pkg/globals/api"
	"gorm.io/gorm/clause"
)

func SeederFunc(model interface{}, message string, customFuncs ...func()) {
	if err := api.App.DB.Gorm.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(model); err.Error != nil {
		api.App.Logger.Panic().Err(err.Error).Msg("")
	}

	api.App.DB.Gorm.Save(model)

	for _, customFunc := range customFuncs {
		customFunc()
	}

	api.App.Logger.Info().Msg(message + " has seeded successfully!")
}
