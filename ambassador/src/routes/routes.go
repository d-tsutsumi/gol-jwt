package routes

import (
	"ambassador/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Post("/api/admin/register", controllers.Register)
}
