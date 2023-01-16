package entities

import (
	"time"
)

type UserRoles struct {
	UserID      int64     `json:"user_id"`
	RoleID      int64     `json:"role_id"`
	VerifyToken string    `json:"verify_token"`
	VerifyAt    time.Time `json:"verify_at"`
}

func (ur UserRoles) TableName() string {
	return "user_roles"
}

func NewUserRoles(user_id, role_id int64, verify_token string, verify_at time.Time) *UserRoles {
	return &UserRoles{
		UserID:      user_id,
		RoleID:      role_id,
		VerifyToken: verify_token,
		VerifyAt:    verify_at,
	}
}
