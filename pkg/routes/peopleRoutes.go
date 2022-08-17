package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitPeopleRoutes(people fiber.Router) {
	people.Get("/:id", utils.NotImplemented)          // Get people By people ID
	people.Get("/:id/anime", utils.NotImplemented)    // Get Anime By people ID
	people.Get("/:id/manga", utils.NotImplemented)    // Get Manga By people ID
	people.Get("/:id/voices", utils.NotImplemented)   // Get Voice Actor By people ID
	people.Get("/:id/pictures", utils.NotImplemented) // Get pictures by people ID
	people.Get("/", utils.NotImplemented)             // Search peoples
}
