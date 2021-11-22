package seeds

import (
	"fmt"

	"github.com/ofsahof/library-management/pkg/database"
)

type Seeder interface {
	Seed()
	ReturnModel() interface{}
}

func SeedModels(seeder ...Seeder) {
	for _, v := range seeder {
		var count int64 = 0
		database.DB.Gorm.Model(v.ReturnModel()).Count(&count)

		if count > 0 {
			v.Seed()
		} else {
			fmt.Print("=====> WARN: Table has seeded already. Skipping!")
		}
	}
}
