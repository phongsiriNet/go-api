package database

import (
	product_model "go-api/pkg/model"
	model "go-api/pkg/model/examplemodel"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {

	models := []interface{}{
		&model.ExampleModel{},
		&product_model.Product{},
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}
