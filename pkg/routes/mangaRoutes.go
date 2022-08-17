package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitMangaRoutes(manga fiber.Router) {
	manga.Get("/:id", utils.NotImplemented)                 // Get manga by ID
	manga.Get("/:id/characters", utils.NotImplemented)      // Get manga Characters
	manga.Get("/:id/news", utils.NotImplemented)            // Get manga news
	manga.Get("/:id/forum", utils.NotImplemented)           // Get forum by Manga ID
	manga.Get("/:id/pictures", utils.NotImplemented)        // Get pictures by Manga ID
	manga.Get("/:id/statistics", utils.NotImplemented)      // Get Statistics by Manga ID
	manga.Get("/:id/moreinfo", utils.NotImplemented)        // Get more information about Manga by Manga ID
	manga.Get("/:id/recommendations", utils.NotImplemented) // Get recommendations by Manga ID
	manga.Get("/:id/userupdates", utils.NotImplemented)     // Get user Updates by Manga ID
	manga.Get("/:id/reviews", utils.NotImplemented)         // Get reviews about Manga by Manga ID
	manga.Get("/:id/relations", utils.NotImplemented)       // Get relations by Manga ID
	manga.Get("/:id/external", utils.NotImplemented)        // Get external by Manga ID
	manga.Get("/", utils.NotImplemented)                    // Search Manga
}
