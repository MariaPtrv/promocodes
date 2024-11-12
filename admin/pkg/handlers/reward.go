package handler

import (
	t "admin/pkg"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) NewReward(c echo.Context) error {
	r := new(t.Reward)

	if err := c.Bind(r); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json")
	}

	reward := t.Reward{
		Title: r.Title,
		Desc:  r.Desc,
	}

	rj, _ := json.Marshal(reward)

	log.Printf("handler-reward: CreateReward \nreward: %s\n", string(rj))
	rewardId, err := h.services.Reward.CreateReward(reward)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, "NewReward reward: "+strconv.Itoa(rewardId))
}


func (h *Handler) DeleteReward(c echo.Context) error {
	r := new(t.Reward)

	if err := c.Bind(r); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json")
	}

	reward := t.Reward{
		Title: r.Title,
	}

	log.Printf("handler-reward: DeleteReward \nreward: %s\n", reward.Title)
	err := h.services.Reward.DeleteReward(reward)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, "Deleted reward "+reward.Title)
}
