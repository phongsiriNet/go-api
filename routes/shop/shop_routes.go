package shop

import (
	"go-api/handler"
	"go-api/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(api fiber.Router, proHandler handler.IProductHandler) {
	api.Get("/All", proHandler.GetProducts)
	api.Get("/:id", proHandler.GetProduct)
	api.Post("/createProduct", middleware.RoleMiddleware("admin"), proHandler.CreateProduct)
	api.Put("/updateProduct/:id", middleware.RoleMiddleware("admin"), proHandler.UpdateProduct)
	api.Delete("/:id", middleware.RoleMiddleware("admin"), proHandler.DeleteProduct)
}
func UserRoutes(api fiber.Router, userHangler handler.IUserHandler) {
	apiGroup := api.Group("/user")
	apiGroup.Post("Register", userHangler.Register)
	apiGroup.Post("Login", userHangler.Login)
}
