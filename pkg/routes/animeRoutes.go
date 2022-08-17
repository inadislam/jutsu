package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/app/controllers/anime"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitAnimeRoutes(anime fiber.Router) {
	anime.Get("/:id", controllers.AnimeHandler)               // Get anime by anime ID
	anime.Get("/:id/characters", controllers.GetCharacter)        // Get anime character by anime ID
	anime.Get("/:id/staff", controllers.GetStaff)             // Get anime character by anime ID
	anime.Get("/:id/episodes", utils.NotImplemented)          // Get anime epsisodes by anime ID
	anime.Get("/:id/episodes/:episode", utils.NotImplemented) // Get anime episode by anime ID
	anime.Get("/:id/news", utils.NotImplemented)              // Get anime news by anime ID
	anime.Get("/:id/forum", utils.NotImplemented)             // Get anime forum by anime ID
	anime.Get("/:id/videos", utils.NotImplemented)            // Get anime videos by anime ID
	anime.Get("/:id/pictures", utils.NotImplemented)          // Get anime pictures by anime ID
	anime.Get("/:id/statistics", utils.NotImplemented)        // Get anime statistics by anime ID
	anime.Get("/:id/moreinfo", utils.NotImplemented)          // Get more info by anime ID
	anime.Get("/:id/recommendations", utils.NotImplemented)   // Get recommended anime by anime ID
	anime.Get("/:id/userupdates", utils.NotImplemented)       // Get user update by anime ID
	anime.Get("/:id/reviews", utils.NotImplemented)           // Get anime Reviews by anime ID
	anime.Get("/:id/relations", utils.NotImplemented)         // Get relations by anime ID
	anime.Get("/:id/themes", utils.NotImplemented)            // Get anime theme by anime ID
	anime.Get("/:id/external", utils.NotImplemented)          // Get external info by anime ID
	anime.Get("/", utils.NotImplemented)                      // anime search
}
