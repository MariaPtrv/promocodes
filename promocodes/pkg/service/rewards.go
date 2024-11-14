package service

import (
	t "promocodes"
	"promocodes/pkg/repository"
)

type RewardsService struct {
	repo repository.Rewards
}

func NewRewardsService(repo repository.Rewards) *RewardsService {
	return &RewardsService{repo: repo}
}

func (s *RewardsService) NewRewardsRecord(r t.RewardsRecord) error {
	return s.repo.NewRewardsRecord(r)
}

func (s *RewardsService) GetRewardsRecordByUserId(r t.RewardsRecord) (t.RewardsRecord, error) {
	return s.repo.GetRewardsRecordByUserId(r)
}
