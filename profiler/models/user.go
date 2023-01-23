package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User Model
type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	FirstName   string    `json:"first_name" gorm:"not null" validate:"required"`
	LastName    string    `json:"last_name" gorm:"not null" validate:"required"`
	Phone       string    `json:"phone" gorm:"not null" validate:"required"`
	Email       string    `json:"email" gorm:"unique;not null" validate:"required"`
	Password    string    `json:"password" gorm:"not null" validate:"required"`
	CreditCards []Card
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
