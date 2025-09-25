package database

import (
	shop_model "go-api/pkg/model"
	model "go-api/pkg/model/examplemodel"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {

	models := []interface{}{
		&model.ExampleModel{},
		&shop_model.Product{},
		&shop_model.User{},
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}
