package model

type FinancialReport struct {
	Transactions []Transaction `json:"transactions" bson:"transactions"`
}
