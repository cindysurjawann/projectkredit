package generateSkala

import (
	"errors"
	"fmt"
	"kredit/backend/generateCustomer"
	"kredit/backend/model"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type SkalaRepository interface {
	GenerateSkalaRentalTab() ([]model.CustomerDataTab, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GenerateSkalaRentalTab() ([]model.CustomerDataTab, error) {
	//Get data yang “approval_status” = 0
	var CustomerDataTab []model.CustomerDataTab
	rows := r.db.Preload("LoanDataTab").Where("approval_status=?", "0").Find(&CustomerDataTab)
	if rows.Error != nil {
		if errors.Is(rows.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data to proceed not found")
		}
		return nil, rows.Error
	}

	//proses insert ke skala_rental_tab
	for _, data := range CustomerDataTab {
		loanDataResult := data.LoanDataTab
		var counter int64
		var err error
		if counter, err = generateCustomer.ConvertStringtoInt(loanDataResult.LoanPeriod); err != nil {
			return nil, errors.New("gagal convert loan period")
		}
		//inisialisasi data
		principle, interest, endBalanceLastInt, endBalance, rentalInt, dueDate, err := InitializeSkalaRental(loanDataResult)
		if err != nil {
			return nil, errors.New("gagal inisialisasi data skala rental")
		}

		//insert data sebanyak 0-counter
		for i := int64(0); i <= counter; i++ {
			i, principle, interest, endBalanceLastInt, endBalance, rentalInt, dueDate, loanDataResult, data, err = r.InsertSkalaRentalTab(i, counter, principle, interest, endBalanceLastInt, endBalance, rentalInt, dueDate, loanDataResult, data)
			if err != nil {
				return nil, errors.New("gagal insert ke skala rental tab")
			}
			//finalisasi data
			endBalanceLastInt = endBalance
			//update approval status
			if i == counter {
				if err := r.UpdateApprovalStatus(data.Custcode, "1"); err != nil {
					return nil, errors.New("gagal update approval status")
				}
			}
		}

	}

	return CustomerDataTab, nil
}

func (r *repository) UpdateApprovalStatus(custcode string, approvalStatus string) error {
	var CustomerDataTab []model.CustomerDataTab
	update := r.db.Model(&CustomerDataTab).Where("custcode=?", custcode).Update("approval_status", approvalStatus)
	return update.Error
}

func InitializeSkalaRental(loanDataResult model.LoanDataTab) (int64, int64, int64, int64, int64, time.Time, error) {
	//hitung monthly payment jika 0
	var principle, interest, endBalanceLastInt, endBalance, rentalInt int64
	var dueDate time.Time
	var err error
	if endBalanceLastInt, err = ConvertMoneytoInt(loanDataResult.LoanAmount); err != nil {
		return principle, interest, endBalanceLastInt, endBalance, rentalInt, dueDate, errors.New("gagal convert loan amount ke int")
	}
	if rentalInt, err = ConvertMoneytoInt(loanDataResult.MonthlyPayment); err != nil {
		return principle, interest, endBalanceLastInt, endBalance, rentalInt, dueDate, errors.New("gagal convert monthly payment ke int")
	}
	return principle, interest, endBalanceLastInt, endBalance, rentalInt, dueDate, nil
}

func (r *repository) InsertSkalaRentalTab(
	i int64, counter int64, principle int64, interest int64, endBalanceLastInt int64, endBalance int64, rentalInt int64, dueDate time.Time, loanDataResult model.LoanDataTab, data model.CustomerDataTab) (
	int64, int64, int64, int64, int64, int64, time.Time, model.LoanDataTab, model.CustomerDataTab, error) {
	if i == 0 {
		interest = 0
		principle = 0
		endBalance = endBalanceLastInt
		dueDate = data.DrawdownDate
	} else {
		dueDate = data.DrawdownDate.AddDate(0, int(1*i), 0)
		interest = (endBalanceLastInt * int64(loanDataResult.InterestEffective) * 30) / 36000
		principle = rentalInt - interest
		endBalance = endBalanceLastInt - principle
	}
	if i == counter && endBalance < 0 {
		principle = principle + endBalance
		interest = rentalInt - principle
		endBalance = 0
	}
	SkalaRentalTab := model.SkalaRentalTab{
		Counter:    int8(i),
		Custcode:   data.Custcode,
		OsBalance:  strconv.FormatInt(endBalanceLastInt, 10),
		EndBalance: strconv.FormatInt(endBalance, 10),
		DueDate:    dueDate,
		EffRate:    float64(loanDataResult.InterestEffective),
		Rental:     loanDataResult.MonthlyPayment,
		Interest:   strconv.FormatInt(interest, 10),
		Principle:  strconv.FormatInt(principle, 10),
		InputDate:  time.Now(),
		InputBy:    "System",
	}
	insertSkalaRentalTab := r.db.Create(&SkalaRentalTab)
	if insertSkalaRentalTab.Error != nil {
		return i, principle, interest, endBalanceLastInt, endBalance, rentalInt, dueDate, loanDataResult, data, errors.New("gagal insert ke skala rental tab")
	}
	return i, principle, interest, endBalanceLastInt, endBalance, rentalInt, dueDate, loanDataResult, data, nil
}

func ConvertMoneytoInt(money string) (int64, error) {
	res := int64(0)
	var err error
	money = strings.Replace(money, "Rp", "", -1)
	money = strings.Replace(money, ".", "", -1)
	fmt.Println(money)
	if res, err = generateCustomer.ConvertStringtoInt(money); err != nil {
		return 0, errors.New("gagal convert money to int")
	}
	return res, nil
}
