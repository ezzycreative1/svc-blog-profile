package mysql

import (
	"context"
	"time"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	DB             *gorm.DB
	KeyTransaction string
	timeout        time.Duration
}

func NewUserRepository(db *gorm.DB, keyTransaction string, timeout int) *userRepository {
	return &userRepository{
		DB:             db,
		KeyTransaction: keyTransaction,
		timeout:        time.Duration(timeout) * time.Second,
	}
}

func (ur *userRepository) Store(ctx context.Context, input entities.Users) error {
	trx, ok := ctx.Value(ur.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = ur.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, ur.timeout)
	defer cancel()

	query := trx.WithContext(ctxWT).Create(&input)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
