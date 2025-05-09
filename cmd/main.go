package main

import (
	"go-api/cmd/database"
	serverconfig "go-api/config/server_config"
	"go-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] - [${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 : 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	routes.InitRoutes(app, db)

	app.Listen(":" + serverconfig.ServerConfig().PORT)
}
