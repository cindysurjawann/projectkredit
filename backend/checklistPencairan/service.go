package checklistPencairan

import (
	"kredit/backend/model"
	"net/http"
)

type Service interface {
	FindPengajuanByApprovalStatus(approval_status string) ([]model.CustomerDataTab, int, error)
}

type service struct {
	repo PencairanRepository
}

func NewService(repo PencairanRepository) *service {
	return &service{repo}
}

func (s *service) FindPengajuanByApprovalStatus(approval_status string) ([]model.CustomerDataTab, int, error) {
	cdt, err := s.repo.FindPengajuanByApprovalStatus(approval_status)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return cdt, http.StatusOK, nil
}
