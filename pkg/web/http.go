package web

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ResponseFormatter returning formatted JSON response
func ResponseFormatter(ctx *fiber.Ctx, code int, message string, body interface{}, err error) error {
	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"message": message,
			"data":    nil,
			"error":   err.Error(),
		}
	} else {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   nil,
		}
	}

	return ctx.JSON(response)
}

// ResponseErrValidation returning formatted JSON response
func ResponseErrValidation(ctx *fiber.Ctx, message string, errMap map[string]interface{}) error {

	var b strings.Builder
	for k, v := range errMap {
		b.WriteString(fmt.Sprintf("%s : %v, ", k, v))
	}
	errorString := strings.TrimRight(b.String(), ", ")

	response := map[string]interface{}{
		"message":          message,
		"data":             nil,
		"error_validation": errMap,
		"error":            errorString,
	}

	return ctx.JSON(response)
}

// ResponseErrWithFormValidation returning formatted JSON response
func ResponseErrWithFormatValidation(ctx *fiber.Ctx, message string, validation map[string]interface{}) error {
	return ctx.JSON(
		struct {
			Message        string                 `json:"message"`
			FormValidation map[string]interface{} `json:"form_validation"`
		}{Message: message, FormValidation: validation},
	)
}

func ResponseErrValidationWithCode(ctx *fiber.Ctx, message string, errMap map[string]interface{}, code int) error {

	var msg string

	if len(errMap) == 0 {
		msg = message
	} else {
		for _, value := range errMap {
			msg = fmt.Sprintf("%s", value)
			break
		}
	}

	response := map[string]interface{}{
		"message":          msg,
		"data":             nil,
		"error_validation": errMap,
	}

	return ctx.JSON(response)
}

// ResponseFormatter returning formatted JSON response with meta
func ResponseFormatterWithMeta(ctx *fiber.Ctx, code int, message string, body interface{}, meta interface{}, err error) error {
	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   err.Error(),
			"meta":    meta,
		}
	} else {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   nil,
			"meta":    meta,
		}
	}

	return ctx.JSON(response)
}

func ResponseErrValidationWithDefaultMessage(ctx *fiber.Ctx, message string, errMap map[string]interface{}, code int) error {
	var msg string

	if len(errMap) == 0 {
		msg = message
	} else {
		for _, value := range errMap {
			msg = value.(string)
			break
		}
	}

	response := map[string]interface{}{
		"message":          msg,
		"error_validation": errMap,
	}

	return ctx.JSON(response)
}

func ResponseFormaterWithDefaultEmptyObject(ctx *fiber.Ctx, message, errMessage string, data, meta interface{}, errMap map[string]interface{}, code int) error {

	var status bool
	if data == nil || len(errMap) == 0 {
		status = true
	}

	if len(errMap) == 0 {
		errMap = map[string]interface{}{}
	} else {
		var b strings.Builder
		for k, v := range errMap {
			b.WriteString(fmt.Sprintf("%s : %v, ", k, v))
		}
		errMessage = strings.TrimRight(b.String(), ", ")
	}

	// make empty object instead null value
	if data == nil {
		data = make(map[string]string, 0)
	}

	response := map[string]interface{}{
		"status":           status,
		"msg":              message,
		"message":          message,
		"data":             data,
		"error_validation": errMap,
		"error":            errMessage,
	}

	if meta != nil {
		response["meta"] = meta
	}

	return ctx.JSON(response)
}
