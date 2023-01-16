package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/dtos"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/ports"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mid"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mlog"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mvalidator"
	"github.com/ezzycreative1/svc-blog-profile/pkg/web"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UseCaseUsers ports.IUserUsecase
	Validator    mvalidator.Validator
	Logger       mlog.Logger
	Cfg          config.Group
}

func NewUserHandler(
	usecaseUsers ports.IUserUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) UserHandler {
	return UserHandler{
		UseCaseUsers: usecaseUsers,
		Validator:    validator,
		Logger:       logger,
		Cfg:          config,
	}
}

//Handler User
func (uh *UserHandler) Register(ctx *fiber.Ctx) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Context(), requestID)

	var payload dtos.RegisterRequestBody
	if err := ctx.BodyParser(&payload); err != nil {
		uh.Logger.WarnT(requestID, fmt.Sprintf("register payload: %s", err.Error()), mlog.Any("payload", payload))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// log payload
	uh.Logger.InfoT(requestID, "register request", mlog.Any("payload", payload))

	// Validate slice payload
	mapErr, err := uh.Validator.Struct(payload)
	if err != nil {
		uh.Logger.WarnT(requestID, fmt.Sprintf("Bad Request: %s", err.Error()))
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	if err := uh.UseCaseUsers.Register(userCtx, payload); err != nil {
		uh.Logger.ErrorT(requestID, "error register data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", nil, nil)
}

func (uh *UserHandler) Login(ctx *fiber.Ctx) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Context(), requestID)

	var payload dtos.LoginRequestBody
	if err := ctx.BodyParser(&payload); err != nil {
		uh.Logger.WarnT(requestID, fmt.Sprintf("login payload: %s", err.Error()), mlog.Any("payload", payload))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// log payload
	uh.Logger.InfoT(requestID, "login request", mlog.Any("payload", payload))

	// Validate slice payload
	mapErr, err := uh.Validator.Struct(payload)
	if err != nil {
		uh.Logger.WarnT(requestID, fmt.Sprintf("Bad Request: %s", err.Error()))
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	result, err := uh.UseCaseUsers.Login(userCtx, payload)
	if err != nil {
		uh.Logger.ErrorT(requestID, "error login data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", result, nil)
}

func (uh *UserHandler) Refresh(ctx *fiber.Ctx) error {
	return web.ResponseFormatter(ctx, http.StatusOK, "Success", nil, nil)
}

func (uh *UserHandler) Logout(ctx *fiber.Ctx) error {
	return web.ResponseFormatter(ctx, http.StatusOK, "Success", nil, nil)
}

func (uh *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Context(), requestID)

	idP, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		uh.Logger.ErrorT(requestID, "error get data", err)
		return web.ResponseFormatter(ctx, http.StatusNotFound, err.Error(), nil, err)
	}

	id := int64(idP)

	var payload dtos.UpdateUserRequestBody
	if err := ctx.BodyParser(&payload); err != nil {
		uh.Logger.WarnT(requestID, fmt.Sprintf("update user payload: %s", err.Error()), mlog.Any("payload", payload))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// log payload
	uh.Logger.InfoT(requestID, "update user request", mlog.Any("payload", payload))

	// Validate slice payload
	mapErr, err := uh.Validator.Struct(payload)
	if err != nil {
		uh.Logger.WarnT(requestID, fmt.Sprintf("Bad Request: %s", err.Error()))
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	if err := uh.UseCaseUsers.UpdateUser(userCtx, id, payload); err != nil {
		uh.Logger.ErrorT(requestID, "error update user data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", nil, nil)
}

func (uh *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Context(), requestID)

	idP, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		uh.Logger.ErrorT(requestID, "error get data", err)
		return web.ResponseFormatter(ctx, http.StatusNotFound, err.Error(), nil, err)
	}

	id := int64(idP)

	err = uh.UseCaseUsers.DeleteUser(userCtx, id)
	if err != nil {
		uh.Logger.ErrorT(requestID, "error delete data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", nil, nil)
}

func (uh *UserHandler) FetchUser(ctx *fiber.Ctx) error {

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", nil, nil)
}

func (uh *UserHandler) GetUser(ctx *fiber.Ctx) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Context(), requestID)

	idP, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		uh.Logger.ErrorT(requestID, "error get data", err)
		return web.ResponseFormatter(ctx, http.StatusNotFound, err.Error(), nil, err)
	}

	id := int64(idP)

	result, err := uh.UseCaseUsers.GetUserByID(userCtx, id)
	if err != nil {
		uh.Logger.ErrorT(requestID, "error delete data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", result, nil)
}
