package service

import "admin/pkg/repository"

type RewardService struct {
	repo repository.Reward
}

func NewRewardService(repo repository.Reward) *RewardService {
	return &RewardService{repo: repo}
}

func (s *RewardService) CreateReward(title, desc string) (int, error) {
	return s.repo.CreateReward(title, desc)
}

func (s *RewardService) DeleteReward(title, desc string) (int, error) {
	return s.repo.CreateReward(title, desc)
}

func (s *RewardService) GetReward(title, desc string) (int, error) {
	return s.repo.CreateReward(title, desc)
}

func (s *RewardService) UpdateReward(title, desc string) (int, error) {
	return s.repo.CreateReward(title, desc)
}