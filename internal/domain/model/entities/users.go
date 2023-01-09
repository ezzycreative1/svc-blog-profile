package entities

import (
	"time"
)

type Users struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	RoleID      int64     `json:"role_id"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u Users) TableName() string {
	return "users"
}

func NewUsers(id int64, firstname, lastname, email, password string, roleid int64, status int) *Users {
	return &Users{
		ID:        id,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  password,
		RoleID:    roleid,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
