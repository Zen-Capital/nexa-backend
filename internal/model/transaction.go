package model

import "time"

type Transaction struct {
	Type     string    `json:"type" bson:"type"`
	Category string    `json:"category" bson:"category"`
	Value    float64   `json:"value" bson:"value"`
	Date     time.Time `json:"date" bson:"date"`
}
