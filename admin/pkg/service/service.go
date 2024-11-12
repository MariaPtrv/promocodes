package service

import (
	types "admin/pkg"
	"admin/pkg/repository"
)

type Promocode interface {
	CreatePromocode()
	DeletePromocode()
	UpdatePromocode()
	GetPromocode()
}

type Reward interface {
	CreateReward(r types.Reward) (int, error)
	DeleteReward(r types.Reward) error
	GetRewardById(r types.Reward) (types.Reward, error)
}

type Rewards interface {
	GetRewards() ([]types.Reward, error)
}

type Service struct {
	Promocode
	Reward
	Rewards
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Reward: NewRewardService(repos.Reward),
		Rewards: NewRewardsService(repos.Rewards),
	}
}
