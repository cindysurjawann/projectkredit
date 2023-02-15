package checklistPencairan

import (
	"errors"
	"kredit/backend/model"
	"time"

	"gorm.io/gorm"
)

type PencairanRepository interface {
	FindPengajuanByApprovalStatus(approval_status string) ([]model.CustomerDataTab, error)
	FindPengajuanByFilter(approval_status string, branch string, company string, startDate time.Time, endDate time.Time) ([]model.CustomerDataTab, error)
	GetBranchList() ([]model.BranchTab, error)
	GetCompanyList() ([]model.MstCompanyTab, error)
	UpdateApprovalStatus(CustomerDataTabUpdate []model.CustomerDataTab, approval_status string) ([]model.CustomerDataTab, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPengajuanByApprovalStatus(approval_status string) ([]model.CustomerDataTab, error) {
	var CustomerDataTab []model.CustomerDataTab

	cdt := r.db.Preload("LoanDataTab").Order("ppk ASC").Where("approval_status=?", approval_status).Find(&CustomerDataTab)
	if cdt.Error != nil {
		if errors.Is(cdt.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, cdt.Error
	}

	return CustomerDataTab, nil
}

func (r *repository) FindPengajuanByFilter(approval_status string, branch string, company string, startDate time.Time, endDate time.Time) ([]model.CustomerDataTab, error) {
	var CustomerDataTab []model.CustomerDataTab
	var CustomerDataTabFinal []model.CustomerDataTab
	var cdt *gorm.DB

	if company == "All Company" && branch == "0" {
		cdt = r.db.Preload("LoanDataTab").Order("ppk ASC").Where("approval_status=? AND drawdown_date BETWEEN ? AND ?", approval_status, startDate, endDate).Find(&CustomerDataTab)
	} else if company != "All Company" && branch == "0" {
		cdt = r.db.Preload("LoanDataTab").Order("ppk ASC").Where("approval_status=? AND channeling_company =? AND drawdown_date BETWEEN ? AND ?", approval_status, company, startDate, endDate).Find(&CustomerDataTab)
	} else if company == "All Company" && branch != "0" {
		cdt = r.db.Preload("LoanDataTab").Order("ppk DESC").Where("approval_status=? AND drawdown_date BETWEEN ? AND ?", approval_status, startDate, endDate).Find(&CustomerDataTab)
	} else if company != "All Company" && branch != "0" {
		cdt = r.db.Preload("LoanDataTab").Order("ppk DESC").Where("approval_status=? AND channeling_company =? AND drawdown_date BETWEEN ? AND ?", approval_status, company, startDate, endDate).Find(&CustomerDataTab)
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

func (r *repository) GetBranchList() ([]model.BranchTab, error) {
	var BranchTab []model.BranchTab

	bt := r.db.Order("code ASC").Find(&BranchTab)
	if bt.Error != nil {
		return nil, bt.Error
	}

	FirstList := model.BranchTab{
		Code:        "0",
		Description: "All Branch",
	}

	BranchTab = append([]model.BranchTab{FirstList}, BranchTab...)

	return BranchTab, nil
}

func (r *repository) GetCompanyList() ([]model.MstCompanyTab, error) {
	var MstCompanyTab []model.MstCompanyTab

	mct := r.db.Order("company_name ASC").Find(&MstCompanyTab)
	if mct.Error != nil {
		return nil, mct.Error
	}

	FirstList := model.MstCompanyTab{
		CompanyCode: "000",
		CompanyName: "All Company",
	}

	MstCompanyTab = append([]model.MstCompanyTab{FirstList}, MstCompanyTab...)

	return MstCompanyTab, nil
}

func (r *repository) UpdateApprovalStatus(CustomerDataTabUpdate []model.CustomerDataTab, approval_status string) ([]model.CustomerDataTab, error) {
	for _, data := range CustomerDataTabUpdate {
		cdt := r.db.Where("ppk=?", data.PPK).Updates(model.CustomerDataTab{
			ApprovalStatus: approval_status,
		})
		if cdt.Error != nil {
			return nil, cdt.Error
		}
	}
	return CustomerDataTabUpdate, nil
}
