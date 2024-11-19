package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email      string             `json:"email" bson:"email"`
	Username   string             `json:"username" bson:"username"`
	Password   string             `json:"password" bson:"password"`
	Account    Account            `json:"account" bson:"account"`
	IsActive   bool               `json:"is_active" bson:"is_active"`
	DateJoined time.Time          `json:"date_joined" bson:"date_joined"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`                     // Última atualização
	DeletedAt  *time.Time         `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"` // Exclusão lógica
}
