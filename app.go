package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("kebard1988 OOPS 👋!")
	})

	app.Listen(":3000")
}
