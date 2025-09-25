package service

import (
	"go-api/dto"
	"go-api/pkg/model"

	"github.com/golang-jwt/jwt/v5"
)

type IProductService interface {
	CreateProductSVC(*dto.ProductRequest) (*dto.ProductResponse, error)
	GetProduct(uint) (*dto.ProductResponse, error)
	GetProducts() ([]dto.ProductResponse, error)
	UpdateProductSVC(uint, *dto.ProductRequest) (*dto.ProductResponse, error)
	DeleteProductSVC(uint) error
}

type IUserService interface {
	RegisterSVC(*dto.UserRequest) error
	LoginSVC(*dto.UserRequest) (string, error)
}

type IAuth interface {
	CreateJwt(*model.User) (string, error)
	ValidateJwt(string) (*jwt.Token, error)
}
