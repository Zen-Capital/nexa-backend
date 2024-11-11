package main

import (
	"github.com/Zen-Capital/nexa-backend/database"
	"github.com/Zen-Capital/nexa-backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
