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
	promocode.POST("/new", NewPromocode)
	promocode.GET("/:id", GetPromocode)
	promocode.DELETE("/:id", DeletePromocode)
	promocode.PUT("/:id", UpdatePromocode)

	reward := admin.Group("/reward")
	reward.POST("/new", NewReward)
	reward.GET("/:id", GetReward)
	reward.DELETE("/:id", func(c echo.Context) error {
		return c.String(200, "Delete reward :id")
	})
	reward.PUT("/:id", func(c echo.Context) error {
		return c.String(200, "Update reward :id")
	})
}
