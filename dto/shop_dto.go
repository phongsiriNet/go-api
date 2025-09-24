package dto

type ProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductResponse struct {
	ProductID uint    `json:"product id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserResponse struct {
	Jwt string `json:"jwt"`
}
