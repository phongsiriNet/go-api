package service

import (
	"go-api/dto"
	"go-api/pkg/repository"
)

type productService struct {
	productRepo repository.IProduct
}

func NewProductService(productRepo repository.IProduct) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) CreateProductSVC(product *dto.ProductRequest) (dto.ProductResponse, error) {
	return dto.ProductResponse{}, nil
}

func (s *productService) GetProduct(id uint) (*dto.ProductResponse, error) {
	return nil, nil
}

func (s *productService) GetProducts() ([]dto.ProductResponse, error) {
	return nil, nil
}

func (s *productService) UpdateProductSVC(id uint, updateProduct *dto.ProductRequest) (dto.ProductResponse, error) {
	return dto.ProductResponse{}, nil
}

func (s *productService) DeleteProductSVC(id uint) error {
	return nil
}
