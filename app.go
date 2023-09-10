package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fiber/configs"
	"github.com/fiber/handlers"
	"github.com/fiber/middleware"
	"github.com/fiber/router"
	"github.com/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	en := configs.GetEnv()
	app := fiber.New(fiber.Config{
		AppName:       fmt.Sprintf("Boiler-plate %s", en.Version),
		CaseSensitive: true,
		Prefork:       en.IsPork,
	})
	basePath := fmt.Sprintf("/v%s", strconv.Itoa(en.MajorVersion))
	defaultRouter, newRouter := router.InitRouter(app, ""), router.InitRouter(app, basePath)

	connectConfig()
	setMiddleware(app, basePath)
	setFuncMiddleware(app)

	// healthCheck
	defaultRouter.SetRouter("/", "GET", handlers.HandleHealth)
	setRouter(newRouter, basePath)
	setErrorHandling(newRouter)
	listen(app, en.Port)
}

func connectConfig() {
	// database
	// swagger
}

func setMiddleware(app *fiber.App, basePath string) {
	// recover
	app.Use(recover.New())

	// logger
	utils.LogSetting(app)
}

func setFuncMiddleware(app *fiber.App) {
	middleware.CommonMiddleware(app, "/")
}

func setRouter(router *router.Router, basePath string) {
	router.SetRouter("/hello", "GET", handlers.HandleHello)
	router.SetRouter("/bye", "GET", handlers.HandleBye)
}

func setErrorHandling(router *router.Router) {
	router.ErrorRoute("/*", handlers.HandleNotFound)
}

func listen(app *fiber.App, port string) {
	utils.Info(fmt.Sprintf("connect Server is Running : %s", port))
	log.Fatalln(app.Listen(":3000"))
}
