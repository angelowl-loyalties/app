package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	MinSpend float64 `json:"min_spend"`
	Start time.Time `json:"start_date`
	Duration time.Time `json:"end_date"`
	RewardProgram string `json:"reward_program"`
	RewardAmount int `json:"reward_amount"` 
	MCC int `json:"mcc"`
	Merchant string `json:"merchant"`
}

func (campaign *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	campaign.ID = uuid.New()

	// To remove when we decide how we taking in date params
	campaign.Start = time.Now()
	campaign.Duration = time.Now()
	return
}
