package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitProducerRoutes(producer fiber.Router) {
	producer.Get("/", utils.NotImplemented) // Get Producers By Producers ID
}
