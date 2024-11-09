package service

import "admin/pkg/repository"

type Promocode interface {
}

type Reward interface {
}

type Service struct {
	Promocode
	Reward
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
