package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CardType struct {
	CardType    string    `json:"card_type" gorm:"primaryKey" binding:"required"`
	Name        string    `json:"name" gorm:"not null" binding:"required"`
	ID          uuid.UUID `json:"id" gorm:"unique;type:uuid;<-:create"`
	RewardUnit  string    `json:"reward_unit" gorm:"not null" binding:"required"`
	CardProgram string    `json:"card_program" gorm:"not null" binding:"required"`
	Cards       []Card    // one card type has many cards of that type
}

func (cardType *CardType) BeforeCreate(tx *gorm.DB) (err error) {
	cardType.ID = uuid.New()
	return
}
