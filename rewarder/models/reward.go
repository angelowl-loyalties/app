package models

import (
	"context"
	"log"

	"github.com/gocql/gocql"
)

type Reward struct {
	ID              gocql.UUID `json:"id"`               // cassandra uuid
	CardID          gocql.UUID `json:"card_id"`          // cassandra uuid
	Merchant        string     `json:"merchant"`         // cassandra text
	MCC             int        `json:"mcc"`              // cassandra int
	Currency        string     `json:"currency"`         // cassandra text
	Amount          float64    `json:"amount"`           // cassandra double
	SGDAmount       float64    `json:"sgd_amount"`       // cassandra double
	TransactionID   string     `json:"transaction_id"`   // cassandra text
	TransactionDate string     `json:"transaction_date"` // cassandra text
	CardPAN         string     `json:"card_pan"`         // cassandra text
	CardType        string     `json:"card_type"`        // cassandra text
	RewardAmount    float64    `json:"reward_amount"`    // cassandra double
	Remarks         string     `json:"remarks"`          // cassandra text
}

func RewardCreate(reward Reward) error {
	ctx := context.Background()

	if err := DB.Query(`INSERT INTO transactions.rewards 
						(id, card_id, merchant, mcc, currency, amount, sgd_amount, transaction_id, transaction_date, card_pan, card_type, reward_amount, remarks) 
						VALUES 
						(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		reward.ID, reward.CardID, reward.Merchant, reward.MCC, reward.Currency, reward.Amount, reward.SGDAmount, reward.TransactionID,
		reward.TransactionDate, reward.CardPAN, reward.CardType, reward.RewardAmount, reward.Remarks).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
