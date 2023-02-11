package models

import (
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
	var id, cardId gocql.UUID
	var merchant, currency, tranId, tranDate, cardPan, cardType, remarks string
	var mcc int
	var amount, sgdAmount, rewardAmount float64

	scanner := DB.Query("SELECT * FROM transactions.rewards").Iter().Scanner()

	for scanner.Next() {
		err := scanner.Scan(&id, &amount, &cardId, &cardPan, &cardType, &currency, &mcc,
			&merchant, &remarks, &rewardAmount, &sgdAmount, &tranDate, &tranId)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

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

func RewardGetByCardID(reqCardId string) (rewards []Reward, _ error) {
	var id, cardId gocql.UUID
	var merchant, currency, tranId, tranDate, cardPan, cardType, remarks string
	var mcc int
	var amount, sgdAmount, rewardAmount float64

	query := DB.Query("SELECT * FROM transactions.rewards WHERE card_id = " + reqCardId)
	log.Println(query.String())
	scanner := query.Iter().Scanner()

	for scanner.Next() {
		err := scanner.Scan(&id, &amount, &cardId, &cardPan, &cardType, &currency, &mcc,
			&merchant, &remarks, &rewardAmount, &sgdAmount, &tranDate, &tranId)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

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