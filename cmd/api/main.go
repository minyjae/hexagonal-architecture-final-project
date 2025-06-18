package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	if err := app.Listen(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
