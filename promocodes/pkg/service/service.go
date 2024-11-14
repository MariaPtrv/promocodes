package service

import (
	t "promocodes"
	"promocodes/pkg/repository"
)

type Promocode interface {
	GetPromocode(p t.Promocode) (t.Promocode, error)
	UpdatePromocode(p t.Promocode) (int, error)
}

type Reward interface {
	GetRewardById(r t.Reward) (t.Reward, error)
}

type Rewards interface {
	NewRewardsRecord(r t.RewardsRecord) error
	GetRewardsRecordByUserId(record t.RewardsRecord) (t.RewardsRecord, error)
}

type Service struct {
	Promocode
	Reward
	Rewards
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Reward:    NewRewardService(repos.Reward),
		Promocode: NewPromocodeService(repos.Promocode),
		Rewards:   NewRewardsService(repos.Rewards),
	}
}
