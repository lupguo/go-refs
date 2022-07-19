package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)
	api.Post("/account-log", createAccountLog)
	app.Listen(":3000")
}