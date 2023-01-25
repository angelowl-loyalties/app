package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User Model
type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	FirstName   string    `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName    string    `json:"last_name" gorm:"type:varchar(255);not null"`
	Phone       string    `json:"phone" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"-" gorm:"not null"`
	Role        string    `gorm:"type:varchar(255);not null"`
	CreditCards []Card    // one user has many credit cards
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewUser struct {
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Phone           string `json:"phone" binding:"required,e164"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
