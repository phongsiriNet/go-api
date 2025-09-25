package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `gorm:"unique;not null;default:Unknown"`
	Price float64
	Stock int
}
type User struct {
	ID       uint   `gorm:"colcumn:user_id;autoIncreasment"`
	Email    string `gorm:"not null;unique"`
	Password string
	Role     string
}
