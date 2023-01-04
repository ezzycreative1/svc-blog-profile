package ports

import (
	"context"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/entities"
)

type IUserRepository interface {
	Store(ctx context.Context, input entities.Users) error
}
