package entities

type UserRoles struct {
	UserID      int64  `json:"user_id"`
	RoleID      int64  `json:"role_id"`
	CreatedAt   int64  `json:"created_at"`
	VerifyToken string `json:"verify_token"`
	VerifyAt    int64  `json:"verify_at"`
}

func (ur UserRoles) TableName() string {
	return "user_roles"
}

func NewUserRoles(user_id, role_id, verify_at, created_at int64, verify_token string) *UserRoles {
	return &UserRoles{
		UserID:      user_id,
		RoleID:      role_id,
		CreatedAt:   created_at,
		VerifyToken: verify_token,
		VerifyAt:    verify_at,
	}
}
