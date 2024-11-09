package types

import "time"

type Promocode struct {
	Id          int       `json:"id,omitempty"`
	Promocode   string    `json:"promocode,omitempty"`
	RewardId    int       `json:"reward_id,omitempty"`
	ExpiresAt   time.Time `json:"expires_at,omitempty"`
	MaxUses     int       `json:"max_uses,omitempty"`
	UsesRemains int       `json:"uses_remains,omitempty"`
}


type Reward struct {
	Id        int    `json:"id,omitempty"`
	Promocode string `json:"promocode,omitempty"`
}
