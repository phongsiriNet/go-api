package routes

import (
	"go-api/handler"
	"go-api/routes/shop"

	"github.com/gofiber/fiber/v2"
)

func InitShopRoutes(app *fiber.App, proHand handler.IProductHandler) {
	api := app.Group("shop")

	shop.ProductRoutes(api, proHand)
}
