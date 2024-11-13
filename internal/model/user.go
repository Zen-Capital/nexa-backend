package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Username   string             `json:"username" bson:"username"`
	Password   string             `json:"password" bson:"password"`
	Email      string             `json:"email" bson:"email"`
	Account    Account            `json:"account" bson:"account"`
	IsActive   bool               `json:"is_active" bson:"is_active"`
	DateJoined time.Time          `json:"date_joined" bson:"date_joined"`
}
