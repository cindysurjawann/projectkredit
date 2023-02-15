package generateSkala

import (
	"kredit/backend/model"
	"net/http"
)

type Service interface {
	GenerateSkalaRentalTab() ([]model.CustomerDataTab, int, error)
}

type service struct {
	repo SkalaRepository
}

func NewService(repo SkalaRepository) *service {
	return &service{repo}
}

func (s *service) GenerateSkalaRentalTab() ([]model.CustomerDataTab, int, error) {
	skala, err := s.repo.GenerateSkalaRentalTab()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return skala, http.StatusOK, nil
}
