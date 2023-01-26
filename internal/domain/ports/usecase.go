package ports

import (
	"context"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/dtos"
)

type IUserUsecase interface {
	Register(ctx context.Context, input dtos.RegisterRequestBody) (*dtos.RegisterResponseBody, error)
	Login(ctx context.Context, input dtos.LoginRequestBody) (*dtos.LoginResponseBody, error)
	UpdateUser(ctx context.Context, id int64, input dtos.UpdateUserRequestBody) error
	DeleteUser(ctx context.Context, id int64) error
	GetUserByID(ctx context.Context, id int64) (*dtos.UserResponseBody, error)
}

type IRoleUsecase interface {
}

type ICategoryUsecase interface {
}

type IArticleUsecase interface {
}

type ICommentUsecase interface {
}

type ITagUsecase interface {
}
