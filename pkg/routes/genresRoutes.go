package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitGenresRoutes(genre fiber.Router) {
	genre.Get("/anime", utils.NotImplemented) // Get anime Genres
	genre.Get("/manga", utils.NotImplemented) // Get manga genres
}
