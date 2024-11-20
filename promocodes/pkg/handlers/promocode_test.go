package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	types "promocodes"
	"promocodes/pkg/service"
	"promocodes/pkg/service/mocks"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type mockPromocode struct {
	Id           int       `json:"id,omitempty"`
	Promocode    string    `json:"promocode,omitempty"`
	Reward_id    int       `json:"reward_id,omitempty"`
	Max_uses     int       `json:"max_uses,omitempty"`
	Remain_uses  int       `json:"remain_uses,omitempty"`
	With_expires bool      `json:"with_expires,omitempty"`
	Expires      time.Time `json:"expires,omitempty"`
}

type mockRewardsRecord struct {
	Id           int       `json:"id,omitempty"`
	Promocode_id int       `json:"promocode_id,omitempty"`
	User_id      int       `json:"user_id,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

type reqBody struct {
	User_id   int    `json:"user_id,omitempty"`
	Promocode string `json:"promocode,omitempty"`
}

func (r reqBody) JSON() []byte {
	j, _ := json.Marshal(r)
	return j
}

var mockGetPromocodeData = []mockPromocode{
	{
		Id:           1,
		Promocode:    "TEST_PROMOCODE",
		Reward_id:    1,
		Max_uses:     2,
		Remain_uses:  2,
		With_expires: false,
	},
	{
		Id:           2,
		Promocode:    "TEST_PROMOCODE_2",
		Reward_id:    1,
		Max_uses:     10,
		Remain_uses:  0,
		With_expires: false,
	},
	{
		Id:           3,
		Promocode:    "TEST_PROMOCODE_3",
		Reward_id:    1,
		With_expires: true,
		Expires:      time.Now().Add(time.Hour),
		Max_uses:     10,
		Remain_uses:  1,
	},
	{
		Id:           4,
		Promocode:    "TEST_PROMOCODE_4",
		Reward_id:    1,
		With_expires: true,
		Expires:      time.Now().Add(-1 * time.Hour),
		Max_uses:     10,
		Remain_uses:  1,
	},
}

var RewardsRecordMock = []mockRewardsRecord{
	{
		Id:           1,
		User_id:      1,
		Timestamp:    time.Now(),
		Promocode_id: 1,
	},
}

type Mock struct {
	Method string
	Args   interface{}
	Resp   interface{}
	Err    error
}

func TestHandler_UsePromocode(t *testing.T) {
	promocode := "TEST_PROMOCODE"
	tests := []struct {
		name                         string
		wantErr                      bool
		body                         reqBody
		err                          error
		mockGetPromocode             bool
		mockGetRewardsRecordByUserId bool
		mockGetRewardById            bool
		mockApplyPromocodeAction     bool
		mockGetPromocodeDataIdx      int
	}{
		{
			name:    "test 1: basic case",
			wantErr: false,
			body: reqBody{
				User_id:   1,
				Promocode: promocode,
			},
			err:                          nil,
			mockGetPromocode:             true,
			mockGetRewardsRecordByUserId: true,
			mockGetRewardById:            true,
			mockApplyPromocodeAction:     true,
			mockGetPromocodeDataIdx:      0,
		},
		{
			name:    "test 2: promocode is empty",
			wantErr: true,
			body: reqBody{
				User_id:   1,
				Promocode: "",
			},
			err:                          fmt.Errorf(""),
			mockGetPromocode:             false,
			mockGetRewardsRecordByUserId: false,
			mockGetRewardById:            false,
			mockApplyPromocodeAction:     false,
			mockGetPromocodeDataIdx:      0,
		},
		{
			name:    "test 3: remains 0 uses",
			wantErr: true,
			body: reqBody{
				User_id:   1,
				Promocode: promocode,
			},
			err:                          fmt.Errorf(""),
			mockGetPromocode:             true,
			mockGetRewardsRecordByUserId: false,
			mockGetRewardById:            false,
			mockApplyPromocodeAction:     false,
			mockGetPromocodeDataIdx:      1,
		},
		{
			name:    "test 4: with expires field",
			wantErr: false,
			body: reqBody{
				User_id:   1,
				Promocode: promocode,
			},
			err:                          nil,
			mockGetPromocode:             true,
			mockGetRewardsRecordByUserId: true,
			mockGetRewardById:            true,
			mockApplyPromocodeAction:     true,
			mockGetPromocodeDataIdx:      2,
		},
		{
			name:    "test 5: expired",
			wantErr: false,
			body: reqBody{
				User_id:   1,
				Promocode: promocode,
			},
			err:                          nil,
			mockGetPromocode:             true,
			mockGetRewardsRecordByUserId: false,
			mockGetRewardById:            false,
			mockApplyPromocodeAction:     false,
			mockGetPromocodeDataIdx:      3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(tt.body.JSON())))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			p := mocks.NewPromocodes(t)

			gp := mockGetPromocodeData[tt.mockGetPromocodeDataIdx]
			var exp *time.Time
			if gp.With_expires {
				exp = &gp.Expires
			} else {
				exp = nil
			}
			getPromocodeReturn := types.Promocode{
				Id:          &gp.Id,
				Promocode:   &gp.Promocode,
				Reward_id:   &gp.Reward_id,
				Expires:     exp,
				Max_uses:    &gp.Max_uses,
				Remain_uses: &gp.Remain_uses,
			}

			if tt.mockGetPromocode {
				p.Mock.On("GetPromocode", types.Promocode{Promocode: &tt.body.Promocode}).Return(getPromocodeReturn, tt.err)
			}

			if tt.mockGetRewardsRecordByUserId {
				p.Mock.On("GetRewardsRecordByUserId", mock.Anything).Return(types.RewardsRecord{
					Id:           nil,
					User_id:      nil,
					Timestamp:    nil,
					Promocode_id: nil,
				}, tt.err)
			}

			if tt.mockGetRewardById {
				p.Mock.On("GetRewardById", mock.Anything).Return(types.Reward{
					Id:          1,
					Title:       "REWARD",
					Description: "Just reward",
				}, tt.err)
			}

			if tt.mockApplyPromocodeAction {
				p.Mock.On("ApplyPromocodeAction", mock.Anything, mock.Anything).Return(tt.err)
			}

			h := &Handler{
				services: &service.Service{
					Promocodes: p,
				},
			}

			err := h.UsePromocode(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.UsePromocode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
