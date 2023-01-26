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
	RoleID      int64  `json:"role_id"`
	IsActive    int    `json:"is_active"`
	VerifyToken string `json:"verify_token"`
	VerifyAt    int64  `json:"verify_at"`
	CreatedAt   int64  `json:"created_at"`
	CreatedBy   int64  `json:"created_by"`
	UpdatedAt   int64  `json:"updated_at"`
	UpdatedBy   int64  `json:"updated_by"`
}

func (u Users) TableName() string {
	return "users"
}

func NewUsers(id, role_id, created_by, updated_by, verify_at int64, firstname, lastname, email, password, phone_number, verify_token string, isActive int) *Users {
	return &Users{
		ID:          id,
		FirstName:   firstname,
		LastName:    lastname,
		Email:       email,
		Password:    password,
		PhoneNumber: phone_number,
		RoleID:      role_id,
		IsActive:    isActive,
		VerifyToken: verify_token,
		VerifyAt:    verify_at,
		CreatedAt:   time.Now().UnixNano(),
		CreatedBy:   created_by,
		UpdatedAt:   time.Now().UnixNano(),
		UpdatedBy:   updated_by,
	}
}
