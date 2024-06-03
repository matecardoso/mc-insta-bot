package routes

import (
	"mc-insta-bot/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", controllers.Login)
}
