package usecase

import (
	"context"
	"time"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/dtos"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/entities"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/ports"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mid"
)

type userUseCase struct {
	Repo ports.IUserRepository
}

func NewUserUsecase(repo ports.IUserRepository) *userUseCase {
	return &userUseCase{
		Repo: repo,
	}
}

func (uc *userUseCase) Register(ctx context.Context, input dtos.Register) error {
	password, _ := mid.HashPassword(input.Password)

	data := entities.Users{
		FirstName: input.Firstname,
		LastName:  input.Lastname,
		Email:     input.Email,
		Password:  password,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.Repo.Store(ctx, data); err != nil {
		return err
	}
	return nil
}
