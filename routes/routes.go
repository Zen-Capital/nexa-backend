package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes() {
	app := fiber.New()
	app.Listen(":3000")
}
