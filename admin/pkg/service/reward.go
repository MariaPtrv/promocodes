package service

import (
	t "admin/pkg"
	"admin/pkg/repository"
)

type RewardService struct {
	repo repository.Reward
}

func NewRewardService(repo repository.Reward) *RewardService {
	return &RewardService{repo: repo}
}

func (s *RewardService) CreateReward(r t.Reward) (int, error) {
	return s.repo.CreateReward(r)
}

func (s *RewardService) DeleteReward(r t.Reward) error {
	return s.repo.DeleteReward(r)
}
