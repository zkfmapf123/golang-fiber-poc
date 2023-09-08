package router

import (
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, basePath string) *Router {
	newRouter := app.Group(basePath)

	return &Router{
		r: newRouter,
	}
}

// Helloworld
// @Summary Hello world
// @Description Hello world
// @ID Hello world
// @Accept json
// @Tags get
// @Param id path int true "Item id"
// @Success 200 {object} Item
// @Failure 400 {object} HTTPError
// @Router /api/item{id} [get]
func (router *Router) Helloworld(path string) {
	router.r.Get(path, func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
}

// ByeWorld
// @Summary Bye world
// @Description Bye world
// @ID Bye world
// @Accept json
// @Tags get
// @Param id path int true "Item id"
// @Success 200 {object} Item
// @Failure 400 {object} HTTPError
// @Router /api/item{id} [get]
func (router *Router) ByeWorld(path string) {
	router.r.Get(path, func(c *fiber.Ctx) error {
		return c.SendString("Bye world")
	})
}
