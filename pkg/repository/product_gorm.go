package repository

import (
	"fmt"
	"go-api/pkg/model"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) IProduct {
	return &productRepository{db: database}
}

func (r *productRepository) CreateProduct(product model.Product) (*model.Product, error) {
	result := r.db.Create(&product)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error)
		return nil, result.Error
	}
	fmt.Println(result.Statement.SQL.String())
	return &product, nil
}

func (r *productRepository) GetAll() ([]model.Product, error) {
	products := []model.Product{}
	result := r.db.Order("id").Find(&products)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error)
		return nil, result.Error
	}
	fmt.Println(result.Statement.SQL.String())
	return products, nil
}

func (r *productRepository) GetByID(id uint) (*model.Product, error) {
	product := model.Product{}
	result := r.db.First(&product, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error)
		return nil, result.Error
	}
	fmt.Println(result.Statement.SQL.String())
	return &product, nil
}

func (r *productRepository) UpdateProduct(id uint, product model.Product) (*model.Product, error) {
	result := r.db.Model(&model.Product{}).Where("id=?", id).Updates(map[string]interface{}{
		"Name":  product.Name,
		"Price": product.Price,
		"Stock": product.Stock,
	})
	if result.Error != nil {
		fmt.Println("Error: ", result.Error)
		return nil, result.Error
	}
	fmt.Println(result.Statement.SQL.String())

	updatedProduct := model.Product{}
	r.db.First(&updatedProduct, id)

	return &updatedProduct, nil
}

func (r *productRepository) DeleteProduct(id uint) error {
	result := r.db.Delete(&model.Product{}, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error)
		return result.Error
	}
	fmt.Println(result.Statement.SQL.String())
	return nil
}
