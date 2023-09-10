package handlers

import "github.com/gofiber/fiber/v2"

func HandleHealth(c *fiber.Ctx) error {
	return c.SendString("success")
}

func HandleHello(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "hello world",
	})
}

func HandleBye(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "bye world",
	})
}

func HandleNotFound(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status":  404,
		"message": "Not Found",
	})
}
