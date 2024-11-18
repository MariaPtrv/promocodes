package service

import (
	t "promocodes"
	"promocodes/pkg/repository"
)

type PromocodesService struct {
	repo repository.Promocodes
}

func NewPromocodesService(repo repository.Promocodes) *PromocodesService {
	return &PromocodesService{repo: repo}
}

func (s *PromocodesService) GetPromocode(p t.Promocode) (t.Promocode, error) {
	return s.repo.GetPromocode(p)
}

func (s *PromocodesService) GetRewardById(r t.Reward) (t.Reward, error) {
	return s.repo.GetRewardById(r)
}

func (s *PromocodesService) GetRewardsRecordByUserId(r t.RewardsRecord) (t.RewardsRecord, error) {
	return s.repo.GetRewardsRecordByUserId(r)
}

func (s *PromocodesService) ApplyPromocodeAction(r t.RewardsRecord, p t.Promocode) error {
	return s.repo.ApplyPromocodeAction(r, p)
}
