package utils

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NotImplemented(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "this handler under working process or undefined",
	})
}

func TrimFileExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

func GetMonthNumber(mon string) int {
	var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	for k, i := range months {
		if i == mon {
			return k + 1
		}
	}
	return 0
}

func GetMalId(url string) int {
	re := regexp.MustCompile("[0-9]+")
	malid, _ := strconv.Atoi(strings.Join(re.FindAllString(url, -1), ""))
	return malid
}
