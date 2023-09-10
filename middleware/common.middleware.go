package middleware

import (
	"github.com/fiber/utils"
	"github.com/gofiber/fiber/v2"
)

func CommonMiddleware(app *fiber.App, path string) {
	app.Use("/", func(c *fiber.Ctx) error {
		utils.Info("fallthrouh")
		return c.Next()
	})
}
