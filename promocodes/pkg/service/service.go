package service

import (
	t "promocodes"
	"promocodes/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Promocodes interface {
	GetPromocode(t.Promocode) (t.Promocode, error)
	UpdatePromocode(t.Promocode) (int, error)
	GetRewardById(t.Reward) (t.Reward, error)
	NewRewardsRecord(t.RewardsRecord) error
	GetRewardsRecordByUserId(t.RewardsRecord) (t.RewardsRecord, error)
}

type Service struct {
	Promocodes
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Promocodes: NewPromocodesService(repos.Promocodes),
	}
}
