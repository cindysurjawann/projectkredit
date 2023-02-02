package drawdownReport

import (
	"kredit/backend/model"
	"net/http"
	"time"
)

type Service interface {
	GetDrawdownReport() ([]model.CustomerDataTab, int, error)
	GetDrawdownReportByFilter(branch string, company string, startDate time.Time, endDate time.Time) ([]model.CustomerDataTab, int, error)
}

type service struct {
	repo PencairanRepository
}

func NewService(repo PencairanRepository) *service {
	return &service{repo}
}

func (s *service) GetDrawdownReport() ([]model.CustomerDataTab, int, error) {
	cdt, err := s.repo.GetDrawdownReport()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return cdt, http.StatusOK, nil
}

func (s *service) GetDrawdownReportByFilter(branch string, company string, startDate time.Time, endDate time.Time) ([]model.CustomerDataTab, int, error) {
	cdt, err := s.repo.GetDrawdownReportByFilter(branch, company, startDate, endDate)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return cdt, http.StatusOK, nil
}
