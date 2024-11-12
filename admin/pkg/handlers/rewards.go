package handler

import (
	t "admin/pkg"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRewards(c echo.Context) error {
	r := new(t.Reward)

	if err := c.Bind(r); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json: "+err.Error())
	}

	log.Printf("handler-reward: GetRewards\n")

	rwrd, err := h.services.Rewards.GetRewards()
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, rwrd)
}
