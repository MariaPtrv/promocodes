package types

import "time"

type Promocode struct {
	Id          *int       `json:"id,omitempty"`
	Promocode   *string    `json:"promocode,omitempty"`
	Reward_id   *int       `json:"reward_id,omitempty"`
	Expires     *time.Time `json:"expires,omitempty"`
	Max_uses    *int       `json:"max_uses,omitempty"`
	Remain_uses *int       `json:"remain_uses,omitempty"`
}

type Reward struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
