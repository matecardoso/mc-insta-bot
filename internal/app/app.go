package app

import (
	"mc-insta-bot/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()

	routes.Setup(app)

	return app
}
