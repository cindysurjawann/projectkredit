package drawdownReport

import (
	"errors"
	"kredit/backend/model"
	"time"

	"gorm.io/gorm"
)

type PencairanRepository interface {
	GetDrawdownReport() ([]model.CustomerDataTab, error)
	GetDrawdownReportByFilter(branch string, company string, startDate time.Time, endDate time.Time) ([]model.CustomerDataTab, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetDrawdownReport() ([]model.CustomerDataTab, error) {
	var CustomerDataTab []model.CustomerDataTab

	cdt := r.db.Preload("LoanDataTab").Order("ppk ASC").Where("approval_status IN ('0','1')").Find(&CustomerDataTab)
	if cdt.Error != nil {
		if errors.Is(cdt.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, cdt.Error
	}

	return CustomerDataTab, nil
}

func (r *repository) GetDrawdownReportByFilter(branch string, company string, startDate time.Time, endDate time.Time) ([]model.CustomerDataTab, error) {
	var CustomerDataTab []model.CustomerDataTab
	var CustomerDataTabFinal []model.CustomerDataTab
	var cdt *gorm.DB

	if company == "All Company" && branch == "0" {
		cdt = r.db.Preload("LoanDataTab").Order("ppk ASC").Where("approval_status IN ('0','1') AND drawdown_date BETWEEN ? AND ?", startDate, endDate).Find(&CustomerDataTab)
	} else if company != "All Company" && branch == "0" {
		cdt = r.db.Preload("LoanDataTab").Order("ppk ASC").Where("approval_status IN ('0','1') AND channeling_company =? AND drawdown_date BETWEEN ? AND ?", company, startDate, endDate).Find(&CustomerDataTab)
	} else if company == "All Company" && branch != "0" {
		cdt = r.db.Preload("LoanDataTab").Order("ppk DESC").Where("approval_status IN ('0','1') AND drawdown_date BETWEEN ? AND ?", startDate, endDate).Find(&CustomerDataTab)
	} else if company != "All Company" && branch != "0" {
		cdt = r.db.Preload("LoanDataTab").Order("ppk DESC").Where("approval_status IN ('0','1') AND channeling_company =? AND drawdown_date BETWEEN ? AND ?", company, startDate, endDate).Find(&CustomerDataTab)
	}
	if cdt.Error != nil {
		if errors.Is(cdt.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, cdt.Error
	}

	if branch != "0" {
		for index, data := range CustomerDataTab {
			if data.LoanDataTab.Branch == branch {
				CustomerDataTabFinal = append([]model.CustomerDataTab{CustomerDataTab[index]}, CustomerDataTabFinal...)
			}
		}

		return CustomerDataTabFinal, nil
	}

	return CustomerDataTab, nil
}
