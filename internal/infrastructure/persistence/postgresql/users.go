package postgresql

import (
	"context"
	"errors"
	"time"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/entities"
	"github.com/ezzycreative1/svc-blog-profile/pkg/errs"
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

func (ur *userRepository) StoreUser(ctx context.Context, input entities.Users) error {
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

func (ur *userRepository) GetUserById(ctx context.Context, UserID int64) (*entities.Users, error) {
	trx, ok := ctx.Value(ur.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = ur.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, ur.timeout)
	defer cancel()

	var userdata entities.Users
	query := trx.WithContext(ctxWT).First(&userdata, UserID)
	if query.Error == gorm.ErrRecordNotFound {
		return nil, errs.ErrNotFound
	}
	if query.Error != nil {
		return nil, query.Error
	}

	return &userdata, nil
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, Email string) (*entities.Users, error) {
	trx, ok := ctx.Value(ur.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = ur.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, ur.timeout)
	defer cancel()

	var userdata entities.Users
	query := trx.WithContext(ctxWT).Where("email = ?", Email).First(&userdata)
	if query.Error == gorm.ErrRecordNotFound {
		return nil, errs.ErrNotFound
	}
	if query.Error != nil {
		return nil, query.Error
	}

	return &userdata, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, input entities.Users) error {
	trx, ok := ctx.Value(ur.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = ur.DB
	}
	if input.ID == 0 {
		return errors.New("users to update must have id")
	}

	ctxWT, cancel := context.WithTimeout(ctx, ur.timeout)
	defer cancel()

	query := trx.WithContext(ctxWT).Save(&input)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, input *entities.Users, id int64) error {
	trx, ok := ctx.Value(ur.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = ur.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, ur.timeout)
	defer cancel()

	query := trx.WithContext(ctxWT).Delete(&input, id)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
