package service

import (
	t "admin/pkg"
	"admin/pkg/repository"
)

type RewardsService struct {
	repo repository.Rewards
}

func NewRewardsService(repo repository.Rewards) *RewardsService {
	return &RewardsService{repo: repo}
}

func (s *RewardsService) GetRewards() ([]t.Reward, error) {
	return s.repo.GetRewards()
}
