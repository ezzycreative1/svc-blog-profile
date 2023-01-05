package ports

import (
	"context"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/dtos"
)

type IUserUsecase interface {
	Register(ctx context.Context, input dtos.RegisterRequestBody) error
	Login(ctx context.Context, input dtos.LoginRequestBody) (*dtos.LoginResponseBody, error)
}
