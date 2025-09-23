package service

import (
	"go-api/dto"
	"go-api/pkg/model"
	"go-api/pkg/repository"
	"log"
)

type productService struct {
	productRepo repository.IProduct
}

func NewProductService(productRepo repository.IProduct) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) CreateProductSVC(product *dto.ProductRequest) (*dto.ProductResponse, error) {
	newProduct := model.Product{Name: product.Name, Price: product.Price, Stock: uint(product.Stock)}
	result, err := s.productRepo.CreateProduct(&newProduct)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	proResponse := dto.ProductResponse{ProductID: result.ID, Name: result.Name, Price: result.Price, Stock: int(result.Stock)}
	return &proResponse, nil
}

func (s *productService) GetProduct(id uint) (*dto.ProductResponse, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	proResponse := dto.ProductResponse{
		ProductID: product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     int(product.Stock),
	}
	return &proResponse, nil
}

func (s *productService) GetProducts() ([]dto.ProductResponse, error) {
	products, err := s.productRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	proResponses := []dto.ProductResponse{}

	for _, product := range products {
		proResponse := dto.ProductResponse{
			ProductID: product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     int(product.Stock),
		}
		proResponses = append(proResponses, proResponse)
	}
	return proResponses, nil
}

func (s *productService) UpdateProductSVC(id uint, updateProduct *dto.ProductRequest) (*dto.ProductResponse, error) {
	updatePro := model.Product{Name: updateProduct.Name, Price: updateProduct.Price, Stock: uint(updateProduct.Stock)}
	result, err := s.productRepo.UpdateProduct(id, &updatePro)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	proResponse := dto.ProductResponse{
		ProductID: result.ID,
		Name:      result.Name,
		Price:     result.Price,
		Stock:     int(result.Stock),
	}
	return &proResponse, nil
}

func (s *productService) DeleteProductSVC(id uint) error {
	err := s.productRepo.DeleteProduct(id)
	if err != nil {
		return err
	}
	log.Println("Delete Sucessful")
	return nil
}
