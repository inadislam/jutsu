package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitUsersRoutes(user fiber.Router) {
	user.Get("/", utils.NotImplemented)                          // User Search
	user.Get("/userbyid", utils.NotImplemented)                  // Get user by ID
	user.Get("/:username", utils.NotImplemented)                 // Get user Profile by Username
	user.Get("/:username/statistics", utils.NotImplemented)      // Get user statistics by username
	user.Get("/:username/favorites", utils.NotImplemented)       // Get favorites by Username
	user.Get("/:username/userupdates", utils.NotImplemented)     // get user updates by Username
	user.Get("/:username/about", utils.NotImplemented)           // get about a user by Username
	user.Get("/:usrename/history", utils.NotImplemented)         // Get user history by Username
	user.Get("/:username/friends", utils.NotImplemented)         // Get user friends list by Username
	user.Get("/:usrename/reviews", utils.NotImplemented)         // Get user reviews by Username
	user.Get("/:username/recommendations", utils.NotImplemented) // Get recommendations by Username
	user.Get("/:username/clubs", utils.NotImplemented)           // Get clubs by Username
}
