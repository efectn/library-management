package utils

import (
	"fmt"

	"github.com/ofsahof/library-management/pkg/database"
)

func SeederFunc(model interface{}, message string) {
	if err := database.DB.Gorm.Create(model); err.Error != nil {
		fmt.Printf("=====> ERROR: %v\n", err.Error)
	}

	fmt.Println("=====> INFO: " + message + " seeded successfully!")
}
