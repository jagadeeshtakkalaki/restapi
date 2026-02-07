package models

import "gorm.io/gorm"

type Example struct {
    gorm.Model
    Name string `json:"name"`
	Email string `json:"email"`
	Uuid string `json:"uuid" gorm:"uniqueIndex"`
	Address string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Age int `json:"age"`
	Profession string `json:"profession"`
	Company string `json:"company"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
