package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitCharacterRoutes(character fiber.Router) {
	character.Get("/:id", utils.NotImplemented)          // Get Character By Character ID
	character.Get("/:id/anime", utils.NotImplemented)    // Get Anime By Character ID
	character.Get("/:id/manga", utils.NotImplemented)    // Get Manga By Character ID
	character.Get("/:id/voices", utils.NotImplemented)   // Get Voice Actor By Character ID
	character.Get("/:id/pictures", utils.NotImplemented) // Get pictures by Character ID
	character.Get("/", utils.NotImplemented)             // Search characters
}
