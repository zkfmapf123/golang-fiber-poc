package main

import (
	"fmt"
	"log"

	"github.com/fiber/src/configs"
	"github.com/fiber/src/utils"
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

	connectConfig()
	setMiddleware(app)
	setFuncMiddleware(app)
	setRouter(app)
	setErrorHandling(app)
	listen(app, en.Port)
}

func connectConfig() {

}

func setMiddleware(app *fiber.App) {
	app.Use(recover.New())
	utils.LogSetting(app)
}

func setFuncMiddleware(app *fiber.App) {

}

func setRouter(app *fiber.App) {

}

func setErrorHandling(app *fiber.App) {

}

func listen(app *fiber.App, port string) {
	utils.Info(fmt.Sprintf("connect Server is Running : %s", port))
	log.Fatalln(app.Listen(fmt.Sprintf(":%s", port)))

}
