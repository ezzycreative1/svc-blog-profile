package handler

import (
	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/ports"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mlog"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mvalidator"
)

type TagHandler struct {
	UseCaseTag ports.ITagUsecase
	Validator  mvalidator.Validator
	Logger     mlog.Logger
	Cfg        config.Group
}

func NewTagHandler(
	usecaseTag ports.ITagUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) TagHandler {
	return TagHandler{
		UseCaseTag: usecaseTag,
		Validator:  validator,
		Logger:     logger,
		Cfg:        config,
	}
}
