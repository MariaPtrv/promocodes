package handler

import (
	t "promocodes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	PromocodeAlreadyUsed = iota
	PromocodeExpired
	PromocodeUsedTooManyTimes
)

type usePromocodeReq struct {
	User_id int `json:"user_id"`
	t.Promocode
}

type usePromocodeResp struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(statusCode int, message string) *echo.HTTPError {
	log.Error(message)
	return echo.NewHTTPError(statusCode, errorResponse{message})
}
