package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	CardPan    string    `json:"card_pan" gorm:"not null" validate:"required;credit_card"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid"`
	CardTypeID uuid.UUID `json:"card_type_id" gorm:"type:uuid"`
	CardType   CardType
}

func (card *Card) BeforeCreate(tx *gorm.DB) (err error) {
	card.ID = uuid.New()
	return
}
