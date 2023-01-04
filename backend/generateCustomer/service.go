package generateCustomer

import (
	"kredit/backend/model"
	"net/http"
)

type Service interface {
	GetStagingCustomer() ([]model.StagingCustomer, int, error)
}

type service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *service {
	return &service{repo}
}

func (s *service) GetStagingCustomer() ([]model.StagingCustomer, int, error) {
	customer, err := s.repo.GetStagingCustomer()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return customer, http.StatusOK, nil
}
