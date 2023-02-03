package models

import (
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
