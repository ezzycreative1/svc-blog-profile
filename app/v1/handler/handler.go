package handler

import (
	"net/http"

	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/ezzycreative1/svc-blog-profile/internal/core/ports"
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

	listAr, err := gh.UseCaseUsers.Register(userCtx)
	if err != nil {
		gh.Logger.ErrorT(requestID, "error fetch data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", listAr, nil)
}

// func (ch *BlogHandler) GetRoleByID(ctx *fiber.Ctx) error {
// 	requestID := mid.GetID(ctx)
// 	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

// 	idP, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		ch.Logger.ErrorT(requestID, "error get data", err)
// 		return web.ResponseFormatter(ctx, http.StatusNotFound, err.Error(), nil, err)
// 	}

// 	id := int64(idP)

// 	art, err := ch.UseCaseRoles.GetRoleByID(userCtx, id)
// 	if err != nil {
// 		ch.Logger.ErrorT(requestID, "error get data", err)
// 		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
// 	}

// 	return web.ResponseFormatter(ctx, http.StatusOK, "success", art, nil)
// }

// func (ch *BlogHandler) StoreRole(ctx *fiber.Ctx) (err error) {
// 	requestID := mid.GetID(ctx)
// 	userCtx := mid.SetIDx(ctx.Context(), requestID)

// 	var payload domain.RoleRequest
// 	if err := ctx.Bind(&payload); err != nil {
// 		ch.Logger.ErrorT(requestID, "role store payload", err, mlog.Any("payload", payload))
// 		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
// 	}

// 	mapErr, err := ch.Validator.Struct(payload)
// 	if err != nil {
// 		ch.Logger.ErrorT(requestID, "Bad Request", err)
// 		return web.ResponseErrValidation(ctx, "bad request", mapErr)
// 	}

// 	err = ch.UseCaseRoles.StoreRole(userCtx, &payload)
// 	if err != nil {
// 		ch.Logger.ErrorT(requestID, "error store data", err)
// 		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
// 	}

// 	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
// }

// func (ch *BlogHandler) UpdateRole(ctx *fiber.Ctx) (err error) {
// 	requestID := mid.GetID(ctx)
// 	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

// 	idP, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		ch.Logger.ErrorT(requestID, "error get data", err)
// 		return web.ResponseFormatter(ctx, http.StatusNotFound, err.Error(), nil, err)
// 	}

// 	id := int64(idP)

// 	var payload domain.RoleRequest
// 	if err := ctx.Bind(&payload); err != nil {
// 		ch.Logger.ErrorT(requestID, "role store payload", err, mlog.Any("payload", payload))
// 		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
// 	}

// 	mapErr, err := ch.Validator.Struct(payload)
// 	if err != nil {
// 		ch.Logger.ErrorT(requestID, "Bad Request", err)
// 		return web.ResponseErrValidation(ctx, "bad request", mapErr)
// 	}

// 	err = ch.UseCaseRoles.UpdateRole(userCtx, id, &payload)
// 	if err != nil {
// 		ch.Logger.ErrorT(requestID, "error update data", err)
// 		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
// 	}

// 	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
// }

// func (ch *BlogHandler) DeleteRole(ctx *fiber.Ctx) error {
// 	requestID := mid.GetID(c)
// 	userCtx := mid.SetIDx(c.Request().Context(), requestID)

// 	idP, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		ch.Logger.ErrorT(requestID, "error delete data", err)
// 		return web.ResponseFormatter(c, http.StatusNotFound, err.Error(), nil, err)
// 	}

// 	id := int64(idP)

// 	err = ch.UseCaseRoles.DeleteRole(userCtx, id)
// 	if err != nil {
// 		ch.Logger.ErrorT(requestID, "error delete data", err)
// 		return web.ResponseFormatter(c, http.StatusBadRequest, err.Error(), nil, err)
// 	}

// 	return web.ResponseFormatter(c, http.StatusNoContent, "success", "", nil)
// }
