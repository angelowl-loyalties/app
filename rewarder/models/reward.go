package models

import (
	"log"

	"github.com/scylladb/gocqlx"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/qb"
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
	CreatedAt       string     `json:"created_at"`       // cassandra text
	CardPAN         string     `json:"card_pan"`         // cassandra text
	CardType        string     `json:"card_type"`        // cassandra text
	RewardAmount    float64    `json:"reward_amount"`    // cassandra double
	Remarks         string     `json:"remarks"`          // cassandra text
}

func RewardCreate(reward Reward) error {
	stmt, names := qb.Insert("transactions.rewards").Columns(
		"id",
		"card_id",
		"merchant",
		"mcc",
		"currency",
		"amount",
		"sgd_amount",
		"transaction_id",
		"transaction_date",
		"created_at",
		"card_pan",
		"card_type",
		"reward_amount",
		"remarks").ToCql()

	q := gocqlx.Query(DB.Query(stmt), names).BindStruct(&reward)

	err := q.ExecRelease()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
