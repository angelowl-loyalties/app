package models

import "github.com/google/uuid"

type Transaction struct {
	ID              uuid.UUID `json:"id"`
	CardID          uuid.UUID `json:"card_id"`
	Merchant        string    `json:"merchant"`
	MCC             int       `json:"mcc"`
	Currency        string    `json:"currency"`
	Amount          float64   `json:"amount"`
	SGDAmount       float64   `json:"sgd_amount"`
	TransactionID   string    `json:"transaction_id"`
	TransactionDate string    `json:"transaction_date"`
	CardPAN         string    `json:"card_pan"`
	CardType        string    `json:"card_type"`
}

var Seed_transaction = Transaction{
	ID:              uuid.MustParse("4aab2f7c-4dd3-4a77-beb8-8582048c9bdb"),
	CardID:          uuid.MustParse("3c0b3d7f-c011-4a7d-b47e-1f7c03a8ca53"),
	Merchant:        "Acme Inc.",
	MCC:             5912,
	Currency:        "SGD",
	Amount:          100.00,
	SGDAmount:       142.40,
	TransactionID:   "1234abcd",
	TransactionDate: "10-02-23",
	CardPAN:         "1234567890123456",
	CardType:        "Visa",
}
