package models

import (
	"github.com/gocql/gocql"
)

type Transaction struct {
	ID              gocql.UUID `json:"ID"` //cassandra uuid
	CardID          gocql.UUID `json:"CardID"` //cassandra uuid
	Merchant        string     `json:"Merchant"` //cassandra text
	MCC             int        `json:"MCC"` //cassandra int
	Currency        string     `json:"Currency"` //cassandra text
	Amount          float64    `json:"Amount"` //cassandra double
	SGDAmount       float64    `json:"SGDAmount"` //cassandra double
	TransactionID   string     `json:"TransactionID"` //cassandra text
	TransactionDate string     `json:"TransactionDate"` //cassandra text
	CardPAN         string     `json:"CardPAN"` //cassandra text
	CardType        string     `json:"CardType"` //cassandra text
}
