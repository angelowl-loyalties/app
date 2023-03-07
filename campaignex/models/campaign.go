package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	Name     string    `json:"name" gorm:"type:varchar(255)" binding:"required,min=1"`
	MinSpend float64   `json:"min_spend" binding:"gte=0"`
	// TODO: Add back the gte to add validation after time now
	Start              time.Time `json:"start_date" binding:"required"`             // should be later than time.Now()
	End                time.Time `json:"end_date" binding:"required,gtfield=Start"` // should be later than Start
	RewardProgram      string    `json:"reward_program" gorm:"type:varchar(255)" binding:"required,min=1"`
	RewardAmount       int       `json:"reward_amount" binding:"required,gt=0"`
	MCC                int       `json:"mcc" binding:"required,gte=1,lte=9999"`
	Merchant           string    `json:"merchant" gorm:"type:varchar(255)" binding:"required,min=1"`
	IsBaseReward       bool      `json:"base_reward"`
	ForForeignCurrency bool      `json:"foreign_currency"`
}

func (campaign *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	campaign.ID = uuid.New()
	return
}

func CampaignGetAll() (campaigns []Campaign, err error) {
	err = DB.Find(&campaigns).Error

	return campaigns, err
}

func CampaignGetById(uuid string) (campaign *Campaign, err error) {
	err = DB.Where("id = ?", uuid).First(&campaign).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return campaign, err
}

func CampaignSave(updatedCampaign *Campaign) (campaign *Campaign, err error) {
	err = DB.Save(&updatedCampaign).Error

	return campaign, err
}

func CampaignCreate(campaign *Campaign) (_ *Campaign, err error) {
	err = DB.Create(&campaign).Error

	return campaign, err
}

func CampaignDelete(campaign *Campaign) (_ *Campaign, err error) {
	err = DB.Delete(&campaign).Error

	return campaign, err
}
