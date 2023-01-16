package handler

import (
	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/ports"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mlog"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mvalidator"
)

type RoleHandler struct {
	UseCaseRoles ports.IRoleUsecase
	Validator    mvalidator.Validator
	Logger       mlog.Logger
	Cfg          config.Group
}

func NewRoleHandler(
	usecaseRoles ports.IRoleUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) RoleHandler {
	return RoleHandler{
		UseCaseRoles: usecaseRoles,
		Validator:    validator,
		Logger:       logger,
		Cfg:          config,
	}
}
