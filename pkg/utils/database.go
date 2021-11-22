package utils

import (
	"fmt"

	"github.com/efectn/library-management/pkg/app"
)

func SeederFunc(model interface{}, message string) {
	if err := app.App.DB.Gorm.Create(model); err.Error != nil {
		fmt.Printf("=====> ERROR: %v\n", err.Error)
	}

	fmt.Println("=====> INFO: " + message + " seeded successfully!")
}
