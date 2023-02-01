package checklistPencairan

import (
	"kredit/backend/model"
	"net/http"
	"time"
)

type Service interface {
	FindPengajuanByApprovalStatus(approval_status string) ([]model.CustomerDataTab, int, error)
	FindPengajuanByFilter(approval_status string, branch string, company string, startDate time.Time, endDate time.Time) ([]model.CustomerDataTab, int, error)
	GetBranchList() ([]model.BranchTab, int, error)
	GetCompanyList() ([]model.MstCompanyTab, int, error)
	UpdateApprovalStatus(CustomerDataTabUpdate []model.CustomerDataTab, approval_status string) ([]model.CustomerDataTab, int, error)
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

func (s *service) FindPengajuanByFilter(approval_status string, branch string, company string, startDate time.Time, endDate time.Time) ([]model.CustomerDataTab, int, error) {
	cdt, err := s.repo.FindPengajuanByFilter(approval_status, branch, company, startDate, endDate)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return cdt, http.StatusOK, nil
}

func (s *service) GetBranchList() ([]model.BranchTab, int, error) {
	bt, err := s.repo.GetBranchList()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return bt, http.StatusOK, nil
}

func (s *service) GetCompanyList() ([]model.MstCompanyTab, int, error) {
	mct, err := s.repo.GetCompanyList()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return mct, http.StatusOK, nil
}

func (s *service) UpdateApprovalStatus(CustomerDataTabUpdate []model.CustomerDataTab, approval_status string) ([]model.CustomerDataTab, int, error) {
	cdt, err := s.repo.UpdateApprovalStatus(CustomerDataTabUpdate, approval_status)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return cdt, http.StatusOK, nil
}
