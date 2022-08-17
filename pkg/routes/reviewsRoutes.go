package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitReviewsRoutes(review fiber.Router) {
	review.Get("/anime", utils.NotImplemented) // Get Anime Reviews
	review.Get("/manga", utils.NotImplemented) // Get Manga Reviews
}
