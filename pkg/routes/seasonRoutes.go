package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitSeasonRoutes(season fiber.Router) {
	season.Get("/", utils.NotImplemented)              // Get seasons
	season.Get("/:year/:season", utils.NotImplemented) // Get seasons by Year and Season
	season.Get("/now", utils.NotImplemented)           // Get seasons now
	season.Get("/upcoming", utils.NotImplemented)      // Get upcoming seasons
}
