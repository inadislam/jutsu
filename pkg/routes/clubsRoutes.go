package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitClubRoutes(club fiber.Router) {
	club.Get("/:id", utils.NotImplemented)           // Get clubs by ID
	club.Get("/:id/members", utils.NotImplemented)   // Get members by Clubs ID
	club.Get("/:id/staff", utils.NotImplemented)     // Get staff by Clubs ID
	club.Get("/:id/relations", utils.NotImplemented) // Get staff by Clubs ID
	club.Get("/", utils.NotImplemented)              // Search clubs
}
