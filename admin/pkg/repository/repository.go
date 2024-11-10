package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable      = "user"
	rewardTable    = "reward"
	rewardsTable   = "rewards"
	promocodeTable = "promocode"
)

type Promocode interface {
}

type Reward interface {
	CreateReward(title, desc string) (int, error)
}

type Repository struct {
	Promocode
	Reward
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
