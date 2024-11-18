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

type Promocodes interface {
	GetPromocode(t.Promocode) (t.Promocode, error)
	GetRewardsRecordByUserId(t.RewardsRecord) (t.RewardsRecord, error)
	GetRewardById(t.Reward) (t.Reward, error)	
	ApplyPromocodeAction(t.RewardsRecord, t.Promocode) error
}

type Repository struct {
	Promocodes
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Promocodes: NewPromocodesPostgres(db),
	}
}
