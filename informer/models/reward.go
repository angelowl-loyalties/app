package models

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
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

type Count struct {
	Count int `json:"count"`
}

func RewardGetCount() (count Count, _ error) {
	stmt, _ := qb.Select("angelowl.rewards").CountAll().ToCql()

	err := gocqlx.Select(&count, DB.Query(stmt))
	if err != nil {
		log.Println(err)
		count.Count = -1
		return count, err
	}

	return count, nil
}

func RewardGetByCardID(reqCardId string) (rewards []Reward, _ error) {
	stmt, _ := qb.Select("angelowl.rewards").Where(qb.EqLit("card_id", reqCardId)).AllowFiltering().ToCql()
	err := gocqlx.Select(&rewards, DB.Query(stmt))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rewards, nil
}

func RewardGetCountByCardID(reqCardId string) (count Count, _ error) {
	stmt, _ := qb.Select("angelowl.rewards").CountAll().Where(qb.EqLit("card_id", reqCardId)).AllowFiltering().ToCql()
	err := gocqlx.Select(&count, DB.Query(stmt))
	if err != nil {
		log.Println(err)
		count.Count = -1
		return count, err
	}

	return count, nil
}
