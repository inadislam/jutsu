package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitMegazineRoutes(megazine fiber.Router) {
	megazine.Get("/", utils.NotImplemented) // Get megazines
}
