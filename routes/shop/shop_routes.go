package shop

import (
	"go-api/handler"
	"go-api/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(api fiber.Router, proHandler handler.IProductHandler) {
	api.Get("/all", proHandler.GetProducts)
	api.Get("/:id", proHandler.GetProduct)
	api.Post("/create-product", middleware.RoleMiddleware("admin"), proHandler.CreateProduct)
	api.Put("/update-product/:id", middleware.RoleMiddleware("admin"), proHandler.UpdateProduct)
	api.Delete("/:id", middleware.RoleMiddleware("admin"), proHandler.DeleteProduct)
}
func UserRoutes(api fiber.Router, userHangler handler.IUserHandler) {
	apiGroup := api.Group("/user")
	apiGroup.Post("register", userHangler.Register)
	apiGroup.Post("login", userHangler.Login)
}

// createProduct
