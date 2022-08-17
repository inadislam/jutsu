package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitWatchRoutes(watch fiber.Router) {
	watch.Get("/episodes", utils.NotImplemented)         // Get episodes
	watch.Get("/episodes/popular", utils.NotImplemented) // Get popular episodes
	watch.Get("/promos", utils.NotImplemented)           // Get promos
	watch.Get("/promos/popular", utils.NotImplemented)   // Get popular promos
}
