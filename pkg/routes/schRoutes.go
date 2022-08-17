package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inadislam/jutsu/pkg/utils"
)

func InitSchedulesRoutes(sch fiber.Router) {
	sch.Get("/", utils.NotImplemented) // Get Schedules
}
