package repository

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}
