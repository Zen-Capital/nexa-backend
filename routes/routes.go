package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(client *mongo.Client) {
	app := fiber.New()
	app.Listen(":3000")
}
