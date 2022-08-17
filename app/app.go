package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/routes"
	"github.com/joho/godotenv"
)

func InitApp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}
	fiberApp := fiber.New()
	routes.InitRoutes(fiberApp)
	log.Fatal(fiberApp.Listen(os.Getenv("SERVER_URL")))
}
