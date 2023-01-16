package handler

import (
	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/ports"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mlog"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mvalidator"
)

type ArticleHandler struct {
	UseCaseArticle ports.IArticleUsecase
	Validator      mvalidator.Validator
	Logger         mlog.Logger
	Cfg            config.Group
}

func NewArticleHandler(
	usecaseArticle ports.IArticleUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) ArticleHandler {
	return ArticleHandler{
		UseCaseArticle: usecaseArticle,
		Validator:      validator,
		Logger:         logger,
		Cfg:            config,
	}
}
