package ports

import (
	"context"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/entities"
)

type IUserRepository interface {
	GetUserById(ctx context.Context, UserID int64) (*entities.Users, error)
	GetUserByEmail(ctx context.Context, Email string) (*entities.Users, error)
	StoreUser(ctx context.Context, input entities.Users) error
}
