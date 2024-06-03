package routes

import (
	"mc-insta-bot/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/login", controllers.Login)
}
