package models

import (
	"github.com/gocql/gocql"
)

type Transaction struct {
	ID              gocql.UUID `json:"ID"`
	CardID          gocql.UUID `json:"CardID"`
	Merchant        string     `json:"Merchant"`
	MCC             int        `json:"MCC"`
	Currency        string     `json:"Currency"`
	Amount          float64    `json:"Amount"`
	SGDAmount       float64    `json:"SGDAmount"`
	TransactionID   string     `json:"TransactionID"`
	TransactionDate string     `json:"TransactionDate"`
	CardPAN         string     `json:"CardPAN"`
	CardType        string     `json:"CardType"`
}
