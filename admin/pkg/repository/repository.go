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
	CreatePromocode(p t.Promocode) (int, error)
	GetPromocodeById(p t.Promocode) (t.Promocode, error)
	UpdatePromocode(p t.Promocode) (int, error)
	DeletePromocode(t.Promocode) error
	GetPromocodes() ([]t.Promocode, error)
}

type Reward interface {
	CreateReward(t.Reward) (int, error)
	DeleteReward(t.Reward) error
	GetRewardById(t.Reward) (t.Reward, error)
	GetRewards() ([]t.Reward, error)
}

type Repository struct {
	Promocode
	Reward
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Reward:    NewRewardPostgres(db),
		Promocode: NewPromocodePostgres(db),
	}
}
