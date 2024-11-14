package service

import (
	t "promocodes"
	"promocodes/pkg/repository"
)

type RewardService struct {
	repo repository.Reward
}

func NewRewardService(repo repository.Reward) *RewardService {
	return &RewardService{repo: repo}
}

func (s *RewardService) GetRewardById(r t.Reward) (t.Reward, error) {
	return s.repo.GetRewardById(r)
}
