package checklistPencairan

import (
	"errors"
	"kredit/backend/model"

	"gorm.io/gorm"
)

type PencairanRepository interface {
	FindPengajuanByApprovalStatus(approval_status string) ([]model.CustomerDataTab, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPengajuanByApprovalStatus(approval_status string) ([]model.CustomerDataTab, error) {
	var CustomerDataTab []model.CustomerDataTab

	cdt := r.db.Preload("LoanDataTab").Where("customer_data_tab.approval_status=?", approval_status).Find(&CustomerDataTab)
	if cdt.Error != nil {
		if errors.Is(cdt.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, cdt.Error
	}

	return CustomerDataTab, nil
}
