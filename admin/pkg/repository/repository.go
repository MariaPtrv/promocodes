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

type Rewards interface {
	GetRewards() ([]t.Reward, error)
}

type Reward interface {
	CreateReward(t.Reward) (int, error)
	DeleteReward(t.Reward) error
	GetRewardById(t.Reward) (t.Reward, error)
}

type Repository struct {
	Promocode
	Reward
	Rewards
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Reward:  NewRewardPostgres(db),
		Rewards: NewRewardsPostgres(db),
	}
}
