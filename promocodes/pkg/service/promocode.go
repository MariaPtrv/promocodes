package service

import (
	t "promocodes"
	"promocodes/pkg/repository"
)

type PromocodeService struct {
	repo repository.Promocode
}

func NewPromocodeService(repo repository.Promocode) *PromocodeService {
	return &PromocodeService{repo: repo}
}

func (s *PromocodeService) GetPromocode(p t.Promocode) (t.Promocode, error) {
	return s.repo.GetPromocode(p)
}

func (s *PromocodeService) UpdatePromocode(p t.Promocode) (int, error) {
	return s.repo.UpdatePromocode(p)
}
