package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fiber/src/configs"
	"github.com/fiber/src/middleware"
	"github.com/fiber/src/router"
	"github.com/fiber/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Swagger API
// @version 1.0.0
// @description Hello world
// @contact.name API Support
// @contact.email zkfmapf123@naver.com
// @host localhost:3010
// @BasePath /v1
func main() {

	en := configs.GetEnv()
	app := fiber.New(fiber.Config{
		AppName:       fmt.Sprintf("Boiler-plate %s", en.Version),
		CaseSensitive: true,
		Prefork:       en.IsPork,
	})
	basePath := fmt.Sprintf("/v%s", strconv.Itoa(en.MajorVersion))
	newRouter := router.InitRouter(app, basePath)

	connectConfig()
	setMiddleware(app, basePath)
	setFuncMiddleware(app)
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
	router.Swagger("/swagger/*", basePath)
	router.Helloworld("/hello")
	router.ByeWorld("/bye")
}

func setErrorHandling(router *router.Router) {
	router.ErrorRoute("/*")
}

func listen(app *fiber.App, port string) {
	utils.Info(fmt.Sprintf("connect Server is Running : %s", port))
	log.Fatalln(app.Listen(fmt.Sprintf(":%s", port)))
}
