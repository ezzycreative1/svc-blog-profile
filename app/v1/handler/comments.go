package handler

import (
	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/ports"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mlog"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mvalidator"
)

type CommentHandler struct {
	UseCaseComment ports.ICommentUsecase
	Validator      mvalidator.Validator
	Logger         mlog.Logger
	Cfg            config.Group
}

func NewCommentHandler(
	usecaseComment ports.ICommentUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) CommentHandler {
	return CommentHandler{
		UseCaseComment: usecaseComment,
		Validator:      validator,
		Logger:         logger,
		Cfg:            config,
	}
}
