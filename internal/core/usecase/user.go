package usecase

import (
	"context"

	"github.com/ezzycreative1/svc-blog-profile/core/ports"
)

type userUseCase struct {
	Repo ports.IUserRepository
}

func NewUserUsecase(repo ports.IUserRepository) userUseCase {
	return userUseCase{
		Repo: repo,
	}
}

func (uc *userUseCase) Register(ctx context.Context)
