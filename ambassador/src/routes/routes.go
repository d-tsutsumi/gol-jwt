package routes

import (
	"ambassador/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	api := app.Group("api")
	admin := api.Group("admin")

	admin.Post("/api/admin/register", controllers.Register)
}
