package handler

import (
	t "admin/pkg"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type newPromocodeResp struct {
	Id int `json:"id,omitempty"`
}

func (h *Handler) NewPromocode(c echo.Context) error {
	p := new(t.Promocode)

	if err := c.Bind(p); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json")
	}

	if *p.Max_uses < 1 {
		return newErrorResponse(http.StatusBadRequest, "Invalid Max_uses value. Max_uses must be greater than or equal to 1")
	}

	promocode := t.Promocode{
		Promocode:   p.Promocode,
		Reward_id:   p.Reward_id,
		Max_uses:    p.Max_uses,
		Remain_uses: p.Max_uses,
	}

	var st time.Time
	if *p.Expires != st {
		promocode.Expires = p.Expires

		st = time.Now()
		if st.Unix() >= p.Expires.Unix() {
			return newErrorResponse(http.StatusBadRequest, "Invalid Expires value. Expires must be later than now")
		}
	}

	pj, _ := json.Marshal(promocode)

	log.Printf("handler-promocode: NewPromocode promocode: %s\n", string(pj))

	promocodeId, err := h.services.Promocode.CreatePromocode(promocode)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, newPromocodeResp{
		Id: promocodeId,
	})
}

func (h *Handler) GetPromocode(c echo.Context) error {
	p := new(t.Promocode)

	if err := c.Bind(p); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json: "+err.Error())
	}

	promocode := t.Promocode{
		Id: p.Id,
	}

	log.Printf("handler-promocode: GetPromocode promocode: %d\n", promocode.Id)
	prmcd, err := h.services.Promocode.GetPromocodeById(promocode)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, prmcd)
}

func (h *Handler) DeletePromocode(c echo.Context) error {
	r := new(t.Promocode)

	if err := c.Bind(r); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json")
	}

	promocode := t.Promocode{
		Promocode: r.Promocode,
	}

	log.Printf("handler-promocode: DeletePromocode promocode id: %s\n", promocode.Promocode)
	err := h.services.Promocode.DeletePromocode(promocode)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, struct{}{})
}

func (h *Handler) UpdatePromocode(c echo.Context) error {
	p := new(t.Promocode)

	if err := c.Bind(p); err != nil {
		return newErrorResponse(http.StatusBadRequest, "wrong json")
	}

	if p.Id == nil {
		return newErrorResponse(http.StatusBadRequest, "Invalid 'Id' value. 'Id' must be set")
	}

	if p.Max_uses != nil && *p.Max_uses < 1 {
		return newErrorResponse(http.StatusBadRequest, "Invalid Max_uses value. Max_uses must be greater than or equal to 1")
	}

	promocode := t.Promocode{
		Id:          p.Id,
		Promocode:   p.Promocode,
		Reward_id:   p.Reward_id,
		Max_uses:    p.Max_uses,
		Remain_uses: p.Max_uses,
	}

	var st time.Time
	if p.Expires != nil && *p.Expires != st {
		promocode.Expires = p.Expires

		st = time.Now()
		if st.Unix() >= p.Expires.Unix() {
			return newErrorResponse(http.StatusBadRequest, "Invalid Expires value. Expires must be later than now")
		}
	}

	pj, _ := json.Marshal(promocode)

	log.Printf("handler-promocode: UpdatePromocode promocode: %s\n", string(pj))

	promocodeId, err := h.services.Promocode.UpdatePromocode(promocode)
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, newPromocodeResp{
		Id: promocodeId,
	})
}

func (h *Handler) GetPromocodes(c echo.Context) error {
	log.Printf("handler-promocode: GetPromocodes\n")

	promocodes, err := h.services.Promocode.GetPromocodes()
	if err != nil {
		return newErrorResponse(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, promocodes)
}
