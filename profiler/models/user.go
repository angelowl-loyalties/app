package models

import (
	"errors"
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

type UserInput struct {
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Phone           string `json:"phone" binding:"required,e164"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

func UserGetAll() (users []User, err error) {
	err = DB.Find(&users).Error

	return users, err
}

func UserGetById(uuid string) (user *User, err error) {
	err = DB.Where("id = ?", uuid).Preload("CreditCards").First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}

func UserGetByEmail(email string) (user *User, err error) {
	err = DB.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}

func UserCreate(user *User) (_ *User, err error) {
	err = DB.Omit("CreditCards").Create(&user).Error

	return user, err
}

func UserSave(updatedUser *User) (_ *User, err error) {
	err = DB.Save(&updatedUser).Error

	return updatedUser, err
}

func UserDelete(user *User) (_ *User, err error) {
	err = DB.Delete(&user).Error

	return user, err
}
