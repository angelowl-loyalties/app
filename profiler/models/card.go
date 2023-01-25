package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	CardPan          string    `json:"card_pan" gorm:"unique;not null" binding:"required,credit_card"`
	UserID           uuid.UUID `json:"user_id" gorm:"type:uuid" binding:"required"` // card belongs to one user
	CardTypeCardType string    `json:"card_type" binding:"required"`                // card belongs to one card type
}

func (card *Card) BeforeCreate(tx *gorm.DB) (err error) {
	card.ID = uuid.New()
	return
}
