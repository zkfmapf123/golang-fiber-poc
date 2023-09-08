package router

import "github.com/gofiber/fiber/v2"

func (router *Router) ErrorRoute(path string) {
	router.r.Get(path, func(c *fiber.Ctx) error {
		return c.SendString("Error Handling")
	})
}
