package router

import "github.com/gofiber/fiber/v2"

func (router *Router) ErrorRoute(path string, handler func(c *fiber.Ctx) error) {
	router.r.Get(path, handler)
}
