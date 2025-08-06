package middleware

import (
	apphttp "crypto-price-service/internal/delivery/http/dto"
	apperrors "crypto-price-service/internal/errors"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func ErrorMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("PANIC", slog.Any("error", err))
				c.JSON(
					http.StatusInternalServerError,
					formatErrorResponse(apperrors.NewInternalServerError()),
				)
				c.Abort()
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *apperrors.ApplicationError

			if errors.As(err, &appErr) {
				status := appErr.HttpCode()

				logger.Error("Application error",
					slog.Int("code", appErr.Code()),
					slog.Int("http_code", status),
					slog.String("message", appErr.Message()),
					slog.String("detail", appErr.Detail()),
					slog.Any("cause", appErr.Cause()),
				)

				c.JSON(status, formatErrorResponse(appErr))
			} else {
				logger.Error("Unexpected error",
					slog.Any("error", err),
					slog.String("type", fmt.Sprintf("%T", err)),
				)

				c.JSON(
					http.StatusInternalServerError,
					formatErrorResponse(apperrors.NewInternalServerError()),
				)
			}
			c.Abort()
		}
	}
}

func formatErrorResponse(err *apperrors.ApplicationError) apphttp.ErrorResponse {
	response := apphttp.ErrorResponse{
		Error: apphttp.ErrorBody{
			Code:    err.Code(),
			Message: err.Message(),
			Detail:  err.Detail(),
		},
	}

	// debug-информацию только в dev-режиме
	if gin.Mode() == gin.DebugMode && err.Cause() != nil {
		response.Error.Cause = err.Cause().Error()
	}

	return response
}
