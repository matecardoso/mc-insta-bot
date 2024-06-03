package controllers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	username := os.Getenv("INSTAGRAM_USER")
	password := os.Getenv("INSTAGRAM_PASS")

	log.Printf("Logging in with username: %s and password: %s\n", username, password)

	return c.SendString("Login successful!")
}
