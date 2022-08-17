package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/inadislam/jutsu/app/controllers"
)

func InitRoutes(app *fiber.App) {
	app.Use(cache.New(
		cache.Config{
			Expiration:   24 * time.Hour,
			CacheControl: true,
			CacheHeader:  "X-Cache-Status",
		},
	))
	api := app.Group("/v1")
	api.Get("/", controllers.Home)
	anime := api.Group("/anime")
	InitAnimeRoutes(anime)
	character := api.Group("/characters")
	InitCharacterRoutes(character)
	club := api.Group("/clubs")
	InitClubRoutes(club)
	genre := api.Group("/genres")
	InitGenresRoutes(genre)
	megazine := api.Group("/megazines")
	InitMegazineRoutes(megazine)
	manga := api.Group("/manga")
	InitMangaRoutes(manga)
	people := api.Group("/people")
	InitPeopleRoutes(people)
	producers := api.Group("/producers")
	InitProducerRoutes(producers)
	random := api.Group("/random")
	InitRandomRoutes(random)
	recom := api.Group("/recommendations")
	InitRecommendationsRoutes(recom)
	reviews := api.Group("/reviews")
	InitReviewsRoutes(reviews)
	sch := api.Group("/schedules")
	InitSchedulesRoutes(sch)
	users := api.Group("/users")
	InitUsersRoutes(users)
	season := api.Group("/seasons")
	InitSeasonRoutes(season)
	top := api.Group("/top")
	InitTopRoutes(top)
	watch := api.Group("/watch")
	InitWatchRoutes(watch)
}
