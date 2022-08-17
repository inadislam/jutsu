package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitRecommendationsRoutes(recom fiber.Router) {
	recom.Get("/anime", utils.NotImplemented) // Get recommended Anime
	recom.Get("/manga", utils.NotImplemented) // Get recommended Manga
}
