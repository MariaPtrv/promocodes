package repository

import (
	t "admin/pkg"

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
	CreateReward(t.Reward) (int, error)
	DeleteReward(t.Reward) error
}

type Repository struct {
	Promocode
	Reward
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Reward: NewRewardPostgres(db),
	}
}
