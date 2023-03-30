package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create" binding:"required,uuid4"`
	CardPan          string    `json:"card_pan" gorm:"unique;not null" binding:"required,credit_card"`
	UserID           uuid.UUID `json:"user_id" gorm:"type:uuid" binding:"required"` // card belongs to one user
	CardTypeCardType string    `json:"card_type" binding:"required"`                // card belongs to one card type
}

type CardInput struct {
	ID       string `json:"id" binding:"required,uuid4"`
	CardPan  string `json:"card_pan" binding:"required,credit_card"`
	UserID   string `json:"user_id" binding:"required,uuid4"`
	CardType string `json:"card_type" binding:"required"`
}

//func (card *Card) BeforeCreate(tx *gorm.DB) (err error) {
//	card.ID = uuid.New()
//	return
//}

func CardGetAll() (cards []Card, err error) {
	err = DB.Find(&cards).Error

	return cards, err
}

func CardGetById(uuid string) (card *Card, err error) {
	err = DB.Where("id = ?", uuid).First(&card).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return card, err
}

//func CardGetByPan(pan string) (card *Card, err error) {
//	err = DB.Where("card_pan = ?", pan).First(&card).Error
//
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return nil, nil
//	}
//
//	return card, err
//}

func CardCreate(card *Card) (_ *Card, err error) {
	err = DB.Create(&card).Error

	return card, err
}

func CardDelete(card *Card) (_ *Card, err error) {
	err = DB.Delete(&card).Error

	return card, err
}
