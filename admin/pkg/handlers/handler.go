package handler

import (
	"admin/pkg/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(e *echo.Echo) {

	admin := e.Group("/admin")

	admin.GET("/", func(c echo.Context) error {
		return c.String(200, "Admin here")
	})

	promocode := admin.Group("/promocode")
	promocode.POST("/", NewPromocode)
	promocode.GET("/", GetPromocode)
	promocode.DELETE("/", DeletePromocode)
	promocode.PUT("/", UpdatePromocode)

	reward := admin.Group("/reward")
	reward.POST("/", h.NewReward)
	reward.DELETE("/", h.DeleteReward)
}
