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
type User struct {
	ID       uint   `gorm:"colcumn:user_id;autoIncreasment"`
	Username string `gorm:"not null;unique"`
	Password string
	Role     string
}
