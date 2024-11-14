package handler

import (
	"promocodes/pkg/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(e *echo.Echo) {

	promocodes := e.Group("/promocodes")
	promocode := promocodes.Group("/promocode")
	promocode.POST("/use", h.UsePromocode)

}
