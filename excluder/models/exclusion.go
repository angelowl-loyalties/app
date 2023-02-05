package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Exclusion struct {
	ID  uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	MCC int       `json:"mcc" binding:"required,gte=0,lte=9999"`
}

func (exclusion *Exclusion) BeforeCreate(tx *gorm.DB) (err error) {
	exclusion.ID = uuid.New()
	return
}

func ExclusionGetAll() (exclusions []Exclusion, err error) {
	err = DB.Find(&exclusions).Error

	return exclusions, err
}

func ExclusionGetById(uuid string) (exclusion *Exclusion, err error) {
	err = DB.Where("id = ?", uuid).First(&exclusion).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return exclusion, err
}

func ExclusionSave(exclusion *Exclusion) (_ *Exclusion, err error) {
	err = DB.Save(&exclusion).Error

	return exclusion, err
}

func ExclusionCreate(exclusion *Exclusion) (_ *Exclusion, err error) {
	err = DB.Create(&exclusion).Error

	return exclusion, err
}

func ExclusionDelete(exclusion *Exclusion) (_ *Exclusion, err error) {
	err = DB.Delete(&exclusion).Error

	return exclusion, err
}
