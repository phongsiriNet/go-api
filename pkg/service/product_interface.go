package service

import (
	"go-api/dto"
)

type IProductService interface {
	CreateProductSVC(*dto.ProductRequest) (*dto.ProductResponse, error)
	GetProduct(uint) (*dto.ProductResponse, error)
	GetProducts() ([]dto.ProductResponse, error)
	UpdateProductSVC(uint, *dto.ProductRequest) (*dto.ProductResponse, error)
	DeleteProductSVC(uint) error
}
