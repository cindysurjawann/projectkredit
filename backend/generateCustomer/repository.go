package generateCustomer

import (
	"errors"
	"kredit/backend/model"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetStagingCustomer() ([]model.StagingCustomer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetStagingCustomer() ([]model.StagingCustomer, error) {
	var StagingCustomer []model.StagingCustomer
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")

	rows := r.db.Where("sc_create_date=? AND sc_flag=?", currentDate, "0").Find(&StagingCustomer)
	if rows.Error != nil {
		if errors.Is(rows.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, rows.Error
	}

	for _, data := range StagingCustomer {
		if err := r.ValidateCustomerPPK(data.CustomerPpk); err == nil {
			if err2 := r.UpdateScFlag(data.Id, "8"); err2 != nil {
				return nil, errors.New("gagal update sc flag")
			}
			if err3 := r.InsertStagingError(data, "Duplikasi Customer PPK"); err3 != nil {
				return nil, errors.New("gagal insert ke staging_error")
			}
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		//if validation success
		if err := r.InsertCustomerDataTab(data); err != nil {
			return nil, errors.New("gagal insert customer_data_tab")
		}
		//insertloandatatab
		if err := r.UpdateScFlag(data.Id, "1"); err != nil {
			return nil, errors.New("gagal update sc flag final")
		}
	}

	return StagingCustomer, nil
}

func (r *repository) UpdateScFlag(id int64, scFlag string) error {
	var StagingCustomer []model.StagingCustomer
	update := r.db.Model(&StagingCustomer).Where("id=?", id).Update("sc_flag", scFlag)

	return update.Error
}

func (r *repository) InsertStagingError(data model.StagingCustomer, errorDesc string) error {
	stagingError := model.StagingError{
		SeReff:       data.ScReff,
		SeCreateDate: data.ScCreateDate,
		BranchCode:   data.ScBranchCode,
		Company:      data.ScCompany,
		Ppk:          data.CustomerPpk,
		Name:         data.CustomerName,
		ErrorDesc:    errorDesc,
	}

	insertStagingError := r.db.Create(&stagingError)
	return insertStagingError.Error
}

func (r *repository) InsertCustomerDataTab(data model.StagingCustomer) error {
	birthDate, err := time.Parse("2006-01-02 15:04:05", data.CustomerBirthDate)
	if err != nil {
		return err
	}
	idType, err := strconv.ParseInt(data.CustomerIDType, 10, 8)
	if err != nil {
		return err
	}
	tglPkChanneling, err := time.Parse("2006-01-02", data.LoanTglPkChanneling)
	if err != nil {
		return err
	}
	drawdownDate, err := time.Parse("2006-01-02", data.LoanTglPk)
	if err != nil {
		return err
	}
	customerDataTab := model.CustomerDataTab{
		Custcode:          "C003",
		PPK:               data.CustomerPpk,
		Name:              data.CustomerName,
		Address1:          data.CustomerAddress1,
		Address2:          data.CustomerAddress2,
		City:              data.CustomerCity,
		Zip:               data.CustomerZip,
		BirthPlace:        data.CustomerBirthPlace,
		BirthDate:         birthDate,
		IdType:            int8(idType),
		IdNumber:          data.CustomerIDNumber,
		MobileNo:          data.CustomerMobileNo,
		TglPkChanneling:   tglPkChanneling,
		MotherMaidenName:  data.CustomerMotherMaidenName,
		ApprovalStatus:    "9",
		DrawdownDate:      drawdownDate,
		ChannelingCompany: data.ScCompany,
	}

	insertCustomerDataTab := r.db.Create(&customerDataTab)
	return insertCustomerDataTab.Error
}

func (r *repository) ValidateCustomerPPK(CustomerPpk string) error {
	var CustomerDataTab model.CustomerDataTab
	validasiPpk := r.db.Where("ppk=?", CustomerPpk).First(&CustomerDataTab)

	return validasiPpk.Error
}
