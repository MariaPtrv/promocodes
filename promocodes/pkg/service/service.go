package service

import (
	t "promocodes"
	"promocodes/pkg/repository"
)

type Promocodes interface {
	GetPromocode(t.Promocode) (t.Promocode, error)
	GetRewardById(t.Reward) (t.Reward, error)
	GetRewardsRecordByUserId(t.RewardsRecord) (t.RewardsRecord, error)
	ApplyPromocodeAction(t.RewardsRecord, t.Promocode) error
}

type Service struct {
	Promocodes
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Promocodes: NewPromocodesService(repos.Promocodes),
	}
}
