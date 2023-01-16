package ports

import (
	"context"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/entities"
)

type IUserRepository interface {
	GetUserById(ctx context.Context, UserID int64) (*entities.Users, error)
	GetUserByEmail(ctx context.Context, Email string) (*entities.Users, error)
	StoreUser(ctx context.Context, input entities.Users) error
	UpdateUser(ctx context.Context, input entities.Users) error
	DeleteUser(ctx context.Context, input *entities.Users, id int64) error
}

type IRoleRepository interface {
	GetRoleById(ctx context.Context, RoleID int64) (*entities.Roles, error)
	GetRoleByName(ctx context.Context, Name string) (*entities.Roles, error)
	StoreRole(ctx context.Context, input entities.Roles) error
}
