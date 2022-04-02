package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()

	routes.SetUp(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
	app.Listen(":3000")
}
