package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, headPath string) *Router {
	newRouter := app.Group(fmt.Sprintf("/v%s", headPath))

	return &Router{
		r: newRouter,
	}
}

func (router *Router) Helloworld(path string) {
	router.r.Get(path, func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
}

func (router *Router) ByeWorld(path string) {
	router.r.Get(path, func(c *fiber.Ctx) error {
		return c.SendString("Bye world")
	})
}
