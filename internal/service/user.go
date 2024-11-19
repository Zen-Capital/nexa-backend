package service

import (
	"context"

	"github.com/Zen-Capital/nexa-backend/database"
	"github.com/Zen-Capital/nexa-backend/internal/model"
)

func SaveUser(user model.User) error {
	collection := database.GetCollection("users")
	_, err := collection.InsertOne(context.Background(), user)
	return err
}
