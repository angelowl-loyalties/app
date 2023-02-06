package models

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
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

func RewardGetAll() (rewards []Reward, _ error) {
	var id gocql.UUID
	var cardId gocql.UUID
	var merchant string
	var mcc int
	var currency string
	var amount float64
	var sgdAmount float64
	var tranId string
	var tranDate string
	var cardPan string
	var cardType string
	var remarks string
	var rewardAmount float64

	// temp hard code to work with docker until gocql bug is fixed
	cluster := gocql.NewCluster("cassandra_db")
	cluster.Keyspace = "transactions"

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	// end

	scanner := session.Query("SELECT * FROM transactions.rewards").Iter().Scanner()
	for scanner.Next() {
		err := scanner.Scan(&id, &amount, &cardId, &cardPan, &cardType, &currency, &mcc,
			&merchant, &remarks, &rewardAmount, &sgdAmount, &tranDate, &tranId)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		fmt.Println(id)
		rewards = append(rewards, Reward{
			ID:              id,
			CardID:          cardId,
			Merchant:        merchant,
			MCC:             mcc,
			Currency:        currency,
			Amount:          amount,
			SGDAmount:       sgdAmount,
			TransactionID:   tranId,
			TransactionDate: tranDate,
			CardPAN:         cardPan,
			CardType:        cardType,
			RewardAmount:    rewardAmount,
			Remarks:         remarks,
		})
	}

	return rewards, nil
}
