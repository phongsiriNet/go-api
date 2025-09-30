package routes

import (
	"go-api/handler"
	"go-api/handler/middleware"
	"go-api/pkg/service"
	"go-api/routes/shop"

	"github.com/gofiber/fiber/v2"
)

func InitShopRoutes(app *fiber.App, jwtService service.IAuthService, proHand handler.IProductHandler, userHand handler.IUserHandler) {
	api := app.Group("shop")

	protected := api.Group("/product", middleware.JwtMiddleware(jwtService))
	shop.ProductRoutes(protected, proHand)

	shop.UserRoutes(api, userHand)
}
