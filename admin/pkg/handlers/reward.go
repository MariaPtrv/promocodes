package handler

import (
	t "admin/pkg"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) NewReward(c echo.Context) error {
	r := new(t.Reward)

	if err := c.Bind(r); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json")
	}

	reward := t.Reward{
		Title:       r.Title,
		Description: r.Description,
	}

	rj, _ := json.Marshal(reward)

	log.Printf("handler-reward: CreateReward reward: %s\n", string(rj))
	rewardId, err := h.services.Reward.CreateReward(reward)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, t.Reward{Id: rewardId})
}

func (h *Handler) DeleteReward(c echo.Context) error {
	r := new(t.Reward)

	if err := c.Bind(r); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json")
	}

	reward := t.Reward{
		Title: r.Title,
	}

	log.Printf("handler-reward: DeleteReward reward: %s\n", reward.Title)
	err := h.services.Reward.DeleteReward(reward)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, t.Reward{})
}

func (h *Handler) GetRewardById(c echo.Context) error {
	r := new(t.Reward)

	if err := c.Bind(r); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json: "+err.Error())
	}

	reward := t.Reward{
		Id: r.Id,
	}

	log.Printf("handler-reward: GetReward reward: %d\n", reward.Id)
	rwrd, err := h.services.Reward.GetRewardById(reward)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, rwrd)
}
