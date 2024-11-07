package model

import "time"

type Account struct {
	Balance         float64         `json:"balance" bson:"balance"`
	FinancialReport FinancialReport `json:"financial_report" bson:"financial-report"`
	Investments     []Investment    `json:"investments" bson:"investments"`
	Income          float64         `json:"income" bson:"income"`
	LastUpdated     time.Time       `json:"last_updated" bson:"last_updated"`
}
