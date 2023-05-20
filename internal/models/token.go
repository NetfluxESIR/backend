package models

import "time"

type Token struct {
	Token          string    `json:"token" gorm:"primaryKey"`
	AccountId      string    `json:"account_id"`
	Expire         bool      `json:"expire"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
