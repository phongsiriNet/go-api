package main

import (
	"go-api/cmd/database"
	serverconfig "go-api/config/server_config"
	"go-api/handler"
	"go-api/pkg/repository"
	"go-api/pkg/service"
	"go-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	if err := database.Migration(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	// pro is shortcut of product
	proRepo := repository.NewProductRepository(db)
	proSVC := service.NewProductService(proRepo)
	proHandler := handler.NewProductHandler(proSVC)

	app := fiber.New()
	routes.InitShopRoutes(app, proHandler)

	app.Listen(":" + serverconfig.ServerConfig().PORT)

}
