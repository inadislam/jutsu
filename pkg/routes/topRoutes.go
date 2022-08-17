package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitTopRoutes(top fiber.Router) {
	top.Get("/anime", utils.NotImplemented)      // Get top animes
	top.Get("/manga", utils.NotImplemented)      // Get top Mangas
	top.Get("/people", utils.NotImplemented)     // Get top people's
	top.Get("/characters", utils.NotImplemented) // Get top Characters
	top.Get("/reviews", utils.NotImplemented)    // Get top reviews
}
