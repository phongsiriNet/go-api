package shop

import (
	"go-api/handler"
	"go-api/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(api fiber.Router, proHandler handler.IProductHandler) {
	apiGroup := api.Group("/product")
	apiGroup.Get("/All", proHandler.GetProducts)
	apiGroup.Get("/:id", proHandler.GetProduct)
	apiGroup.Post("/createProduct", middleware.RoleMiddleware("admin"), proHandler.CreateProduct)
	apiGroup.Put("/updateProduct/:id", middleware.RoleMiddleware("admin"), proHandler.UpdateProduct)
	apiGroup.Delete("/:id", middleware.RoleMiddleware("admin"), proHandler.DeleteProduct)
}
func UserRoutes(api fiber.Router, userHangler handler.IUserHandler) {
	apiGroup := api.Group("/user")
	apiGroup.Post("Register", userHangler.Register)
	apiGroup.Post("Login", userHangler.Login)
}
