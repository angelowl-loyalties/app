package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	Name          string    `json:"name" gorm:"type:varchar(255)" binding:"required,min=1"`
	MinSpend      float64   `json:"min_spend" binding:"gte=0"`
	Start         time.Time `json:"start_date" binding:"required,gte"`         // should be later than time.Now()
	End           time.Time `json:"end_date" binding:"required,gtfield=Start"` // should be later than Start
	RewardProgram string    `json:"reward_program" gorm:"type:varchar(255)" binding:"required,min=1"`
	RewardAmount  int       `json:"reward_amount" binding:"required,gt=0"`
	MCC           int       `json:"mcc" binding:"required,gte=0,lte=9999"`
	Merchant      string    `json:"merchant" gorm:"type:varchar(255)" binding:"required,min=1"`
}

func (campaign *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	campaign.ID = uuid.New()
	return
}
