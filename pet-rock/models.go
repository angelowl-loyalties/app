package main

import (
	"github.com/google/uuid"
	"time"
)

//"id","first_name","last_name","phone","email","created_at","updated_at","card_id","card_pan","card_type"

type CSVRecord struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CardID    uuid.UUID `json:"card_id"`
	CardPAN   string    `json:"card_pan"`
	CardType  string    `json:"card_type"`
}

type User struct {
	ID              uuid.UUID `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
}

type Card struct {
	ID       uuid.UUID `json:"id"`
	CardPAN  string    `json:"card_pan"`
	UserID   uuid.UUID `json:"user_id"`
	CardType string    `json:"card_type"`
}

type EmailContent struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type OTPEvent struct {
	Users []EmailContent `json:"users"`
}
