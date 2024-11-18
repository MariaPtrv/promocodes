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

func (s *PromocodesService) UpdatePromocode(p t.Promocode) (int, error) {
	return s.repo.UpdatePromocode(p)
}

func (s *PromocodesService) GetRewardById(r t.Reward) (t.Reward, error) {
	return s.repo.GetRewardById(r)
}


func (s *PromocodesService) NewRewardsRecord(r t.RewardsRecord) error {
	return s.repo.NewRewardsRecord(r)
}

func (s *PromocodesService) GetRewardsRecordByUserId(r t.RewardsRecord) (t.RewardsRecord, error) {
	return s.repo.GetRewardsRecordByUserId(r)
}
