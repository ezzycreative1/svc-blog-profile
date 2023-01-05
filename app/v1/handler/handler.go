package handler

import (
	"fmt"
	"net/http"

	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/dtos"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/ports"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mid"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mlog"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mvalidator"
	"github.com/ezzycreative1/svc-blog-profile/pkg/web"
	"github.com/gofiber/fiber/v2"
)

type BlogHandler struct {
	UseCaseUsers ports.IUserUsecase
	Validator    mvalidator.Validator
	Logger       mlog.Logger
	Cfg          config.Group
}

func NewBlogHandler(
	usecaseUsers ports.IUserUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) BlogHandler {
	return BlogHandler{
		UseCaseUsers: usecaseUsers,
		Validator:    validator,
		Logger:       logger,
		Cfg:          config,
	}
}

//Handler User
func (gh *BlogHandler) Register(ctx *fiber.Ctx) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Context(), requestID)

	var payload dtos.RegisterRequestBody
	if err := ctx.BodyParser(&payload); err != nil {
		gh.Logger.WarnT(requestID, fmt.Sprintf("register payload: %s", err.Error()), mlog.Any("payload", payload))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// log payload
	gh.Logger.InfoT(requestID, "register request", mlog.Any("payload", payload))

	// Validate slice payload
	mapErr, err := gh.Validator.Struct(payload)
	if err != nil {
		gh.Logger.WarnT(requestID, fmt.Sprintf("Bad Request: %s", err.Error()))
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	if err := gh.UseCaseUsers.Register(userCtx, payload); err != nil {
		gh.Logger.ErrorT(requestID, "error register data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", nil, nil)
}

func (gh *BlogHandler) Login(ctx *fiber.Ctx) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Context(), requestID)

	var payload dtos.LoginRequestBody
	if err := ctx.BodyParser(&payload); err != nil {
		gh.Logger.WarnT(requestID, fmt.Sprintf("register payload: %s", err.Error()), mlog.Any("payload", payload))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// log payload
	gh.Logger.InfoT(requestID, "login request", mlog.Any("payload", payload))

	// Validate slice payload
	mapErr, err := gh.Validator.Struct(payload)
	if err != nil {
		gh.Logger.WarnT(requestID, fmt.Sprintf("Bad Request: %s", err.Error()))
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	result, err := gh.UseCaseUsers.Login(userCtx, payload)
	if err != nil {
		gh.Logger.ErrorT(requestID, "error login data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", result, nil)
}

func (gh *BlogHandler) Refresh(ctx *fiber.Ctx) error {
	return web.ResponseFormatter(ctx, http.StatusOK, "Success", nil, nil)
}

func (gh *BlogHandler) Logout(ctx *fiber.Ctx) error {
	return web.ResponseFormatter(ctx, http.StatusOK, "Success", nil, nil)
}
