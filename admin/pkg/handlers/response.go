package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(statusCode int, message string) *echo.HTTPError {
	log.Error(message)
	return echo.NewHTTPError(statusCode, errorResponse{message})
}
