package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CardType struct {
	CardType      string    `json:"card_type" gorm:"primaryKey" binding:"required"`
	Name          string    `json:"name" gorm:"not null" binding:"required"`
	ID            uuid.UUID `json:"id" gorm:"unique;type:uuid;<-:create"`
	RewardUnit    string    `json:"reward_unit" gorm:"not null" binding:"required"`
	RewardProgram string    `json:"reward_program" gorm:"not null" binding:"required"`
	Cards         []Card    `json:"-"` // one card type has many cards of that type
}

func (cardType *CardType) BeforeCreate(tx *gorm.DB) (err error) {
	cardType.ID = uuid.New()
	return
}

func CardTypeGetAll() (cardTypes []CardType, err error) {
	err = DB.Preload("Cards").Find(&cardTypes).Error

	return cardTypes, err
}

func CardTypeGetByType(cardType string) (cardTypeResult *CardType, err error) {
	err = DB.Where("card_type = ?", cardType).Preload("Cards").First(&cardTypeResult).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return cardTypeResult, err
}

func CardTypeCreate(cardType *CardType) (_ *CardType, err error) {
	err = DB.Create(&cardType).Error

	return cardType, err
}

func CardTypeSave(updatedCardType *CardType) (_ *CardType, err error) {
	err = DB.Save(&updatedCardType).Error

	return updatedCardType, err
}

func CardTypeDelete(cardType *CardType) (_ *CardType, err error) {
	err = DB.Delete(&cardType).Error

	return cardType, err
}
