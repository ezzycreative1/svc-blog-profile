package entities

import (
	"time"
)

type Roles struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_atss"`
}

func (u Roles) TableName() string {
	return "roles"
}

func NewRoles(id int64, name string) *Roles {
	return &Roles{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
