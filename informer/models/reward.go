package models

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
)

const YYYYMMDD = "2006-01-02"

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

func RewardGetAll() (rewards []Reward, _ error) {
	stmt, _ := qb.Select("angelowl.rewards").ToCql()

	err := gocqlx.Select(&rewards, DB.Query(stmt))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rewards, nil
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

func RewardGetTodays() (rewards []Reward, _ error) {
	todaysDate := time.Now().Format(YYYYMMDD)
	// Create a string literal with open close quotation marks
	todaysDateLiteral := fmt.Sprintf("'%s'", todaysDate)

	stmt, _ := qb.Select("angelowl.rewards").Where(qb.EqLit("created_at", todaysDateLiteral)).Where(qb.GtLit("reward_amount", "0")).AllowFiltering().ToCql()

	err := gocqlx.Select(&rewards, DB.Query(stmt))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rewards, nil
}
