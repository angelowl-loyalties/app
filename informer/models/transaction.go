package models

import (
	"github.com/gocql/gocql"
)

type Transaction struct {
	ID              gocql.UUID `json:"id"`
	CardID          gocql.UUID `json:"card_id"`
	Merchant        string     `json:"merchant"`
	MCC             int        `json:"mcc"`
	Currency        string     `json:"currency"`
	Amount          float64    `json:"amount"`
	SGDAmount       float64    `json:"sgd_amount"`
	TransactionID   string     `json:"transaction_id"`
	TransactionDate string     `json:"transaction_date"`
	CardPAN         string     `json:"card_pan"`
	CardType        string     `json:"card_type"`
	// add/remove fields for the processed transaction accordingly, above was taken from their example
}
