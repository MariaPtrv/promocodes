package service

import (
	types "admin/pkg"
	"admin/pkg/repository"
)

type Promocode interface {
	CreatePromocode(p types.Promocode) (int, error)
	DeletePromocode(p types.Promocode) error
	UpdatePromocode(p types.Promocode) (int, error)
	GetPromocodeById(p types.Promocode) (types.Promocode, error)
	GetPromocodes() ([]types.Promocode, error)
}

type Reward interface {
	CreateReward(r types.Reward) (int, error)
	DeleteReward(r types.Reward) error
	GetRewardById(r types.Reward) (types.Reward, error)
	GetRewards() ([]types.Reward, error)
}

type Service struct {
	Promocode
	Reward
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Reward:    NewRewardService(repos.Reward),
		Promocode: NewPromocodeService(repos.Promocode),
	}
}
