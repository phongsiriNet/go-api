package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `gorm:"unique;not null;default:Unknown"`
	Price float64
	Stock uint
}
