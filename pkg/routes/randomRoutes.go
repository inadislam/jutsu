package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitRandomRoutes(random fiber.Router) {
	random.Get("/anime", utils.NotImplemented)      // Get random Anime
	random.Get("/manga", utils.NotImplemented)      // Get random manga
	random.Get("/characters", utils.NotImplemented) // Get random Characters
	random.Get("/people", utils.NotImplemented)     // Get random people
	random.Get("/users", utils.NotImplemented)      // Get random users
}
