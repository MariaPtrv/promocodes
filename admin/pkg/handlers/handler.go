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

	promocodes := admin.Group("/promocodes")
	promocodes.GET("/", h.GetPromocodes)

	promocode := promocodes.Group("/promocode")
	promocode.POST("/", h.NewPromocode)
	promocode.GET("/", h.GetPromocode)
	promocode.DELETE("/", h.DeletePromocode)
	promocode.PUT("/", h.UpdatePromocode)

	rewards := admin.Group("/rewards")
	rewards.GET("/", h.GetRewards)

	reward := rewards.Group("/reward")
	reward.POST("/", h.NewReward)
	reward.DELETE("/", h.DeleteReward)
	reward.GET("/", h.GetRewardById)
}
