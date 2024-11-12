package types

import "time"

type Promocode struct {
	Id          int       `json:"id"`
	Promocode   string    `json:"promocode"`
	RewardId    int       `json:"reward_id"`
	Expires     time.Time `json:"expires"`
	MaxUses     int       `json:"max_uses"`
	UsesRemains int       `json:"uses_remains"`
}

type Reward struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Rewards struct {
	Id       int       `json:"id"`
	RewardId string    `json:"title"`
	UserId   string    `json:"user_id"`
	Date     time.Time `json:"date"`
}
