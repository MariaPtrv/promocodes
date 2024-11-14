package handler

import (
	"log"
	"net/http"
	t "promocodes"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UsePromocode(c echo.Context) error {
	p := new(usePromocodeReq)

	if err := c.Bind(p); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json: "+err.Error())
	}

	promocode := t.Promocode{
		Promocode: p.Promocode.Promocode,
	}

	log.Printf("handler-promocode: UsePromocode promocode: %d\n", promocode.Id)

	if promocode.Promocode == nil || len(*promocode.Promocode) < 1 {
		return newErrorResponse(http.StatusBadRequest, "promocode can't be empty")
	}

	prmcd, err := h.services.Promocode.GetPromocode(promocode)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	now := time.Now()

	if prmcd.Expires != nil && prmcd.Expires.Before(now) {
		log.Printf("repository-promocode: UpdatePromocode promocode: %s\n", "promocod has been expired")
		resp := usePromocodeResp{
			Message: "promocod has been expired",
			Status:  PromocodeExpired,
		}
		return c.JSON(200, resp)
	}

	if prmcd.Remain_uses != nil && *prmcd.Remain_uses < 1 {
		log.Printf("repository-promocode: UpdatePromocode promocode: %s\n", "promocode can no longer be applied due to the number of uses")

		resp := usePromocodeResp{
			Message: "promocode can no longer be applied due to the number of uses",
			Status:  PromocodeUsedTooManyTimes,
		}
		return c.JSON(200, resp)
	}

	var userRecordsPromocode t.RewardsRecord

	if prmcd.Id != nil {
		userRecordsPromocode, err = h.services.Rewards.GetRewardsRecordByUserId(t.RewardsRecord{
			User_id:      &p.user_id,
			Promocode_id: prmcd.Id,
		})

		if userRecordsPromocode.Id != nil {
			log.Printf("repository-promocode: UpdatePromocode promocode: %s\n", "promocode has already been used")
			resp := usePromocodeResp{
				Message: "promocode has already been used",
				Status:  PromocodeAlreadyUsed,
			}
			return c.JSON(200, resp)
		}

	}

	reward := t.Reward{
		Id: *prmcd.Reward_id,
	}

	rwrd, err := h.services.Reward.GetRewardById(reward)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	rewardsRecord := t.RewardsRecord{
		Promocode_id: prmcd.Id,
		User_id:      &p.user_id,
		Timestamp:    &now,
	}

	err = h.services.Rewards.NewRewardsRecord(rewardsRecord)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, rwrd)
}