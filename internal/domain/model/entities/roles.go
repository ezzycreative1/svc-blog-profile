package entities

import (
	"time"
)

type Roles struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	IsActive  int    `json:"is_active"`
	CreatedAt int64  `json:"created_at"`
	CreatedBy int64  `json:"created_by"`
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy int64  `json:"updated_by"`
}

func (u Roles) TableName() string {
	return "roles"
}

func NewRoles(id, created_by, updated_by int64, name string, is_active int) *Roles {
	return &Roles{
		ID:        id,
		Name:      name,
		IsActive:  is_active,
		CreatedAt: time.Now().UnixMicro(),
		CreatedBy: created_by,
		UpdatedAt: time.Now().UnixMicro(),
		UpdatedBy: updated_by,
	}
}
