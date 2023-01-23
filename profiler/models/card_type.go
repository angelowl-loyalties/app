package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CardType struct {
	CardType string    `json:"card_type" gorm:"primaryKey"`
	ID       uuid.UUID `json:"id" gorm:"unique;type:uuid;<-:create"`
}

func (cardType *CardType) BeforeCreate(tx *gorm.DB) (err error) {
	cardType.ID = uuid.New()
	return
}
