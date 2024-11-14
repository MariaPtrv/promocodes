package repository

import (
	t "promocodes"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	rewardTable    = "reward"
	rewardsTable   = "rewards"
	promocodeTable = "promocode"
)

type Promocode interface {
	GetPromocode(p t.Promocode) (t.Promocode, error)
	UpdatePromocode(p t.Promocode) (int, error)
}

type Reward interface {
	GetRewardById(t.Reward) (t.Reward, error)
}

type Rewards interface {
	NewRewardsRecord(t.RewardsRecord) error
	GetRewardsRecordByUserId(t.RewardsRecord) (t.RewardsRecord, error)
}

type Repository struct {
	Promocode
	Reward
	Rewards
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Reward:    NewRewardPostgres(db),
		Promocode: NewPromocodePostgres(db),
		Rewards:   NewRewardsPostgres(db),
	}
}
