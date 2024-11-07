package model

import "time"

type Investment struct {
	Name           string    `json: "name" bson:"name"`
	Type           string    `json:"type" bson:"type"`
	Units          int       `json:"units" bson:"units"`
	InvestedValue  float64   `json:"invested_value" bson:"invested_value"`
	AquisitionDate time.Time `json:"aquisition_date" bson:"aquisition_date"`
	SaleDate       time.Time `json:"sale_date" bson:"sale_date"`
	ItIs2YO        bool      `json:"json:"it_is_2yo" bson:"it_is_2yo`
}
