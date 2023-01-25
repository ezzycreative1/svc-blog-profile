package entities

import (
	"time"
)

type Users struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	IsActive    int    `json:"is_active"`
	CreatedAt   int64  `json:"created_at"`
	CreatedBy   int64  `json:"created_by"`
	UpdatedAt   int64  `json:"updated_at"`
	UpdatedBy   int64  `json:"updated_by"`
}

func (u Users) TableName() string {
	return "users"
}

func NewUsers(id, created_by, updated_by int64, firstname, lastname, email, password, phone_number string, isActive int) *Users {
	return &Users{
		ID:          id,
		FirstName:   firstname,
		LastName:    lastname,
		Email:       email,
		Password:    password,
		PhoneNumber: phone_number,
		IsActive:    isActive,
		CreatedAt:   time.Now().UnixNano(),
		CreatedBy:   created_by,
		UpdatedAt:   time.Now().UnixNano(),
		UpdatedBy:   updated_by,
	}
}
