package generateCustomer

import (
	"kredit/backend/model"
	"net/http"
)

type Service interface {
	ValidateStagingCustomer() ([]model.StagingCustomer, int, error)
}

type service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *service {
	return &service{repo}
}

func (s *service) ValidateStagingCustomer() ([]model.StagingCustomer, int, error) {
	customer, err := s.repo.ValidateStagingCustomer()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return customer, http.StatusOK, nil
}
