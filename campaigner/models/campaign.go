package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	MinSpend float64 `json:"minSpend"`
	Start time.Time `json:"startDate`
	Duration time.Time `json:"endDate"`
	RewardProgram string `json:"rewardProgram"`
	RewardAmount int `json:"rewardAmount"` 
	// add other fields here accordingly

}

func (campaign *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	campaign.ID = uuid.New()

	// To remove when we decide how we taking in date params
	campaign.Start = time.Now()
	campaign.Duration = time.Now()
	return
}
