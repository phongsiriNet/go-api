package repository

import "go-api/pkg/model"

type IProduct interface {
	CreateProduct(model.Product) (*model.Product, error)
	GetAll() ([]model.Product, error)
	GetByID(uint) (*model.Product, error)
	UpdateProduct(uint, model.Product) (*model.Product, error)
	DeleteProduct(uint) error
}
