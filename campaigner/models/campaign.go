package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	// add other fields here accordingly

}

func (campaign *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	campaign.ID = uuid.New()
	return
}
