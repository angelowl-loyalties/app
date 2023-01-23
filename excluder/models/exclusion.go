package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Exclusion struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	// add other fields here accordingly

}

func (exclusion *Exclusion) BeforeCreate(tx *gorm.DB) (err error) {
	exclusion.ID = uuid.New()
	return
}
