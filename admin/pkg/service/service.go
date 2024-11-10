package service

import "admin/pkg/repository"

type Promocode interface {
	CreatePromocode()
	DeletePromocode()
	UpdatePromocode()
	GetPromocode()
}

type Reward interface {
	CreateReward(title, desc string) (int, error)
	// DeleteReward()
	// UpdateReward()
	// GetReward()
}

type Service struct {
	Promocode
	Reward
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Reward: NewRewardService(repos.Reward),
	}
}
