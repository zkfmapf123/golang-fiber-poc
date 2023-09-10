package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, basePath string) *Router {
	newRouter := app.Group(basePath)

	return &Router{
		r: newRouter,
	}
}

func (router *Router) SetRouter(path string, method string, handler func(c *fiber.Ctx) error) {

	switch method {
	case "GET":
		router.r.Get(path, handler)
	case "POST":
		router.r.Post(path, handler)
	case "PUT":
		router.r.Put(path, handler)
	case "DELETE":
		router.r.Delete(path, handler)
	default:
		panic(fmt.Sprintf("%s is invalid method\n", method))
	}
}
