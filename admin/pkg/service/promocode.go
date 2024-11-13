package service

import (
	t "admin/pkg"
	"admin/pkg/repository"
)

type PromocodeService struct {
	repo repository.Promocode
}

func NewPromocodeService(repo repository.Promocode) *PromocodeService {
	return &PromocodeService{repo: repo}
}

func (s *PromocodeService) CreatePromocode(p t.Promocode) (int, error) {
	return s.repo.CreatePromocode(p)
}

func (s *PromocodeService) GetPromocodeById(p t.Promocode) (t.Promocode, error) {
	return s.repo.GetPromocodeById(p)
}

func (s *PromocodeService) UpdatePromocode(p t.Promocode) (int, error) {
	return s.repo.UpdatePromocode(p)
}

func (s *PromocodeService) DeletePromocode(p t.Promocode) error {
	return s.repo.DeletePromocode(p)
}

func (s *PromocodeService) GetPromocodes() ([]t.Promocode, error) {
	return s.repo.GetPromocodes()
}
