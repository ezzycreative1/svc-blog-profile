package postgresql

import (
	"context"
	"time"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/entities"
	"github.com/ezzycreative1/svc-blog-profile/pkg/errs"
	"gorm.io/gorm"
)

type roleRepository struct {
	DB             *gorm.DB
	KeyTransaction string
	timeout        time.Duration
}

func NewRoleRepository(db *gorm.DB, keyTransaction string, timeout int) *roleRepository {
	return &roleRepository{
		DB:             db,
		KeyTransaction: keyTransaction,
		timeout:        time.Duration(timeout) * time.Second,
	}
}

func (ur *roleRepository) StoreRole(ctx context.Context, input entities.Roles) error {
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

func (ur *roleRepository) GetRoleById(ctx context.Context, RoleID int64) (*entities.Roles, error) {
	trx, ok := ctx.Value(ur.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = ur.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, ur.timeout)
	defer cancel()

	var roledata entities.Roles
	query := trx.WithContext(ctxWT).First(&roledata, RoleID)
	if query.Error == gorm.ErrRecordNotFound {
		return nil, errs.ErrNotFound
	}
	if query.Error != nil {
		return nil, query.Error
	}

	return &roledata, nil
}

func (ur *roleRepository) GetRoleByName(ctx context.Context, Name string) (*entities.Roles, error) {
	trx, ok := ctx.Value(ur.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = ur.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, ur.timeout)
	defer cancel()

	var roledata entities.Roles
	query := trx.WithContext(ctxWT).Where("name = ?", Name).First(&roledata)
	if query.Error == gorm.ErrRecordNotFound {
		return nil, errs.ErrNotFound
	}
	if query.Error != nil {
		return nil, query.Error
	}

	return &roledata, nil
}
