package shop

import (
	"go-api/handler"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(api fiber.Router, proHandler handler.IProductHandler) {
	api.Group("/product").
		Get("/All", proHandler.GetProducts)
}
