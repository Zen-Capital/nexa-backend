package main

import (
	"github.com/Zen-Capital/nexa-backend/database"
	"github.com/Zen-Capital/nexa-backend/routes"
)

func main() {
	client := database.ConnectDatabase()
	routes.SetupRoutes(client)
}
