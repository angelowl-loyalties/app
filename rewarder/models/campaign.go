package models

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	ID                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	MinSpend           float64   `json:"min_spend"`
	Start              time.Time `json:"start_date"`
	End                time.Time `json:"end_date"`
	RewardProgram      string    `json:"reward_program"`
	RewardAmount       int       `json:"reward_amount"`
	MCC                int       `json:"mcc"`
	Merchant           string    `json:"merchant"`
	IsBaseReward       bool      `json:"base_reward"`
	ForForeignCurrency bool      `json:"foreign_currency"`
}
