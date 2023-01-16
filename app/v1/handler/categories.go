package handler

import (
	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/ports"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mlog"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mvalidator"
)

type CategoryHandler struct {
	UseCaseCategory ports.ICategoryUsecase
	Validator       mvalidator.Validator
	Logger          mlog.Logger
	Cfg             config.Group
}

func NewCategoryHandler(
	usecaseCategory ports.ICategoryUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) CategoryHandler {
	return CategoryHandler{
		UseCaseCategory: usecaseCategory,
		Validator:       validator,
		Logger:          logger,
		Cfg:             config,
	}
}
