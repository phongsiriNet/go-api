package shop

import (
	"go-api/handler"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(api fiber.Router, proHandler handler.IProductHandler) {
	apiGroup := api.Group("/product")
	apiGroup.Get("/All", proHandler.GetProducts)
	apiGroup.Get("/:id", proHandler.GetProduct)
	apiGroup.Post("/createProduct", proHandler.CreateProduct)
	apiGroup.Put("/updateProduct/:id", proHandler.UpdateProduct)
	apiGroup.Delete("/:id", proHandler.DeleteProduct)
}
