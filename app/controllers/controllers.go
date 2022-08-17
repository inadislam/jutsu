package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"author_url":  os.Getenv("AUTHOR_URL"),
		"app_version": os.Getenv("APP_VERSION"),
		"github_url":  os.Getenv("GITHUB_URL"),
	})
}
