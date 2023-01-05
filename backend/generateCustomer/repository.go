package generateCustomer

import (
	"errors"
	"fmt"
	"kredit/backend/model"
	"regexp"
	"strconv"
	"strings"
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
		//1. validate customer ppk
		if err := r.ValidateCustomerPPK(data.CustomerPpk); err == nil {
			if err := r.UpdateWhenValidateFails(data, "Duplikasi Customer PPK"); err != nil {
				return nil, err
			}
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		//2. validate sc company
		if err := r.ValidateCompanyShortName(data.ScCompany); err != nil {
			if err := r.UpdateWhenValidateFails(data, "Company Name tidak terdaftar"); err != nil {
				return nil, err
			}
			continue
		}

		//3. validate branch code
		if err := r.ValidateBranchCode(data.ScBranchCode); err != nil {
			if err := r.UpdateWhenValidateFails(data, "Branch Code tidak terdaftar"); err != nil {
				return nil, err
			}
			continue
		}

		//4. validate loan tgl pk (bulan dan tahun)
		if err := ValidateLoanTglPk(data.LoanTglPk); err != nil {
			if err := r.UpdateWhenValidateFails(data, "Periode Tgl Pk tidak sama dengan Periode sekarang"); err != nil {
				return nil, err
			}
			continue
		}

		//5. validate id number gaboleh kosong jika idtype=1
		if err := ValidateIdNumber(data.CustomerIDType, data.CustomerIDNumber); err != nil {
			if err := r.UpdateWhenValidateFails(data, "ID Number kosong"); err != nil {
				return nil, err
			}
			continue
		}

		//6. validate nama debitur tdk mengandung special karakter
		if err := ValidateNamaDebitur(data.CustomerName); err != nil {
			if err := r.UpdateWhenValidateFails(data, "Nama Debitur mengandung spesial karakter"); err != nil {
				return nil, err
			}
			continue
		}

		//7. validate bpkb gaboleh kosong
		if err := ValidateBpkb(data.VehicleBpkb); err != nil {
			if err := r.UpdateWhenValidateFails(data, "BPKB kosong"); err != nil {
				return nil, err
			}
			continue
		}

		//8. validate stnk gaboleh kosong
		if err := ValidateStnk(data.VehicleStnk); err != nil {
			if err := r.UpdateWhenValidateFails(data, "STNK kosong"); err != nil {
				return nil, err
			}
			continue
		}

		//9. validate vehicle engine no gaboleh kosong
		if err := ValidateEngineNo(data.VehicleEngineNo); err != nil {
			if err := r.UpdateWhenValidateFails(data, "Vehicle Engine No kosong"); err != nil {
				return nil, err
			}
			continue
		}

		//10. validate engine no tidak boleh duplikasi
		if err := r.ValidateDuplicateEngineNo(data.VehicleEngineNo); err == nil {
			if err := r.UpdateWhenValidateFails(data, "Duplikasi Engine No"); err != nil {
				return nil, err
			}
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		//11. validate chasis no gaboleh kosong
		if err := ValidateChasisNo(data.VehicleChasisNo); err != nil {
			if err := r.UpdateWhenValidateFails(data, "Vehicle Chasis No kosong"); err != nil {
				return nil, err
			}
			continue
		}

		//12. validate chasis no tidak boleh duplikasi
		if err := r.ValidateDuplicateChasisNo(data.VehicleChasisNo); err == nil {
			if err := r.UpdateWhenValidateFails(data, "Duplikasi Chasis No"); err != nil {
				return nil, err
			}
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		//if validation success
		if err := r.UpdateWhenValidateSuccess(data); err != nil {
			return nil, err
		}
	}

	return StagingCustomer, nil
}

func (r *repository) UpdateWhenValidateFails(data model.StagingCustomer, errorDesc string) error {
	if err := r.UpdateScFlag(data.Id, "8"); err != nil {
		return errors.New("gagal update sc flag")
	}
	if err := r.InsertStagingError(data, errorDesc); err != nil {
		return errors.New("gagal insert ke staging_error")
	}
	return nil
}

func (r *repository) UpdateWhenValidateSuccess(data model.StagingCustomer) error {
	if err := r.InsertCustomerDataTab(data); err != nil {
		return errors.New("gagal insert customer_data_tab")
	}
	if err := r.InsertLoanDataTab(data); err != nil {
		return errors.New("gagal insert loan_data_tab")
	}
	if err := r.InsertVehicleDataTab(data); err != nil {
		return err
	}
	if err := r.UpdateScFlag(data.Id, "1"); err != nil {
		return errors.New("gagal update sc flag final")
	}
	return nil
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
	idType, err, tglPkChanneling, drawdownDate, birthDate := int64(0), errors.New("initialize error"), time.Now(), time.Now(), time.Now()
	if birthDate, err = ConvertStringtoDateTime(data.CustomerBirthDate); err != nil {
		return err
	}
	if idType, err = ConvertStringtoInt(data.CustomerIDType); err != nil {
		return err
	}
	if tglPkChanneling, err = ConvertStringtoDate(data.LoanTglPkChanneling); err != nil {
		return err
	}
	if drawdownDate, err = ConvertStringtoDate(data.LoanTglPk); err != nil {
		return err
	}
	fmt.Println("InsertCustomerDataTab", idType, err, tglPkChanneling, drawdownDate, birthDate)
	customerDataTab := model.CustomerDataTab{
		Custcode:          "C00" + strconv.FormatInt(data.Id, 10),
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

func (r *repository) InsertLoanDataTab(data model.StagingCustomer) error {
	err, interestflat, interesteffective, effpaymenttype := errors.New("initialize error"), 0.1, 0.1, int64(0)
	if interestflat, err = ConvertStringtoFloat(data.LoanInterestFlatChanneling); err != nil {
		return err
	}
	if interesteffective, err = ConvertStringtoFloat(data.LoanInterestEffectiveChanneling); err != nil {
		return err
	}
	if effpaymenttype, err = ConvertStringtoInt(data.LoanEffectivePaymentType); err != nil {
		return err
	}
	fmt.Println("InsertLoanDataTab", err, interestflat, interesteffective, effpaymenttype)
	loanDataTab := model.LoanDataTab{
		Custcode:             "C00" + strconv.FormatInt(data.Id, 10),
		Branch:               data.ScBranchCode,
		DownPayment:          data.LoanDownPayment,
		LoanAmount:           data.LoanLoanAmountChanneling,
		LoanPeriod:           data.LoanLoanPeriodChanneling,
		InterestFlat:         float32(interestflat),
		InterestEffective:    float32(interesteffective),
		EffectivePaymentType: int8(effpaymenttype),
		MonthlyPayment:       data.LoanMonthlyPaymentChanneling,
		InputtDate:           time.Now(),
		LasttModified:        time.Now(),
		ModifieddBy:          "System",
		InputDate:            time.Now(),
		InputBy:              "System",
		LastModified:         time.Now(),
		ModifiedBy:           "System",
		OTR:                  data.LoanOtr,
		// AdminFee:             data.AdminFee,
		// InterestType:         data.InterestType,
	}

	insertLoanDataTab := r.db.Create(&loanDataTab)

	return insertLoanDataTab.Error
}

func (r *repository) InsertVehicleDataTab(data model.StagingCustomer) error {
	vehicleType, err, vehicleStatus, dealerId, tglStnk, tglBpkb, collateralId := int64(0), errors.New("initialize error"), int64(0), int64(0), time.Now(), time.Now(), int64(0)

	if vehicleType, err = ConvertStringtoInt(data.VehicleType); err != nil {
		return err
	}
	if vehicleStatus, err = ConvertStringtoInt(data.VehicleStatus); err != nil {
		return err
	}
	if dealerId, err = ConvertStringtoInt(data.VehicleDealerID); err != nil {
		return err
	}
	if tglStnk, err = ConvertStringtoDateTime(data.VehicleTglStnk); err != nil {
		return err
	}
	if tglBpkb, err = ConvertStringtoDateTime(data.VehicleTglBpkb); err != nil {
		return err
	}
	if collateralId, err = ConvertStringtoInt(data.CollateralTypeID); err != nil {
		return err
	}

	fmt.Println("InsertVehicleDataTab", vehicleType, err, vehicleStatus, dealerId, tglStnk, tglBpkb, collateralId)
	vehicleDataTab := model.VehicleDataTab{
		Custcode:       "C00" + strconv.FormatInt(data.Id, 10),
		Brand:          int(vehicleType),
		Type:           data.VehicleBrand,
		Year:           data.VehicleYear,
		Jenis:          data.VehicleJenis,
		Status:         int8(vehicleStatus),
		Color:          data.VehicleColor,
		PoliceNo:       data.VehiclePoliceNo,
		EngineNo:       data.VehicleEngineNo,
		ChasisNo:       data.VehicleChasisNo,
		Bpkb:           data.VehicleBpkb,
		Stnk:           data.VehicleStnk,
		DealerID:       int(dealerId),
		InputDate:      time.Now(),
		InputBy:        "System",
		LastModified:   time.Now(),
		ModifiedBy:     "System",
		TglStnk:        tglStnk,
		TglBpkb:        tglBpkb,
		PolisNo:        data.VehiclePoliceNo,
		CollateralID:   collateralId,
		Dealer:         data.VehicleDealer,
		AddressDealer1: data.VehicleAddressDealer1,
		AddressDealer2: data.VehicleAddressDealer2,
		CityDealer:     data.VehicleCityDealer,
		// AgunanLbu: ,
		// KetAgunan: ,
		// TglPolis: ,
		// StnkCity: ,
		// StnkAddress1: ,
		// StnkAddress2: ,
		// RegisterNo: ,
		// Golongan: ,
	}

	insertVehicleDataTab := r.db.Create(&vehicleDataTab)
	return insertVehicleDataTab.Error
}

func (r *repository) ValidateCustomerPPK(CustomerPpk string) error {
	//CUSTOMER_PPK tidak boleh duplikasi pada table Customer_data_tab Field “PPK”
	var CustomerDataTab model.CustomerDataTab
	validasiPpk := r.db.Where("ppk=?", CustomerPpk).First(&CustomerDataTab)

	return validasiPpk.Error
}

func (r *repository) ValidateCompanyShortName(scCompany string) error {
	//SC_COMPANY harus terdaftar di Tabel Mst_Company_Tab Field “Company_Short_Name”
	var MstCompanyTab model.MstCompanyTab
	validasiCompany := r.db.Where("company_short_name=?", scCompany).First(&MstCompanyTab)

	return validasiCompany.Error
}

func (r *repository) ValidateBranchCode(scBranchCode string) error {
	//SC_BRANCH_CODE harus terdaftar di Tabel Branch_Tab Field “Code”
	var BranchTab model.BranchTab
	validasiBranch := r.db.Where("code=?", scBranchCode).First(&BranchTab)

	return validasiBranch.Error
}

func ValidateLoanTglPk(loanTglPk string) error {
	//TGL_PK / DRAWDOWN_DATE tidak boleh berbeda bulan dengan bulan berjalan saat ini
	//contoh: saat ini bulan Januari 2023 maka bulan TGL_PK tidak boleh berbeda dengan bulan Januari 2023
	drawdownDate, err, currentDateTime := time.Now(), errors.New("initialize error"), time.Now()
	if drawdownDate, err = ConvertStringtoDate(loanTglPk); err != nil {
		return err
	}
	fmt.Println("ValidateLoanTglPk", err)

	if drawdownDate.Year() != currentDateTime.Year() {
		return errors.New("drawdown year berbeda")
	}
	if drawdownDate.Month() != currentDateTime.Month() {
		return errors.New("drawdown month berbeda")
	}
	return nil
}

func ValidateIdNumber(customerIDType string, customerIDNumber string) error {
	//Jika “CUSTOMER_ID_TYPE” diisi = 1, maka “CUSTOMER_ID_NUMBER” harus diisi dan tidak boleh kosong
	if customerIDType == "1" {
		if customerIDNumber == "" {
			return errors.New("id number kosong")
		}
		return nil
	}
	return nil
}

func ValidateNamaDebitur(customerName string) error {
	//NAMA Debitur tidak boleh mengandung karakter special
	regex := regexp.MustCompile("^[a-zA-Z ]*$")

	if !regex.MatchString(customerName) {
		return errors.New("nama mengandung karakter special")
	}
	return nil
}

func ValidateBpkb(vehicleBpkb string) error {
	//VEHICLE_BPKB tidak boleh kosong
	if vehicleBpkb == "" {
		return errors.New("bpkb kosong")
	}
	return nil
}

func ValidateStnk(vehicleStnk string) error {
	//	VEHICLE_STNK tidak boleh kosong
	if vehicleStnk == "" {
		return errors.New("stnk kosong")
	}
	return nil
}

func ValidateEngineNo(vehicleEngineNo string) error {
	//	VEHICLE_ENGINE_NO tidak boleh kosong
	if vehicleEngineNo == "" {
		return errors.New("vehicle engine no kosong")
	}
	return nil
}

func (r *repository) ValidateDuplicateEngineNo(vehicleEngineNo string) error {
	//	VEHICLE_ENGINE_NO tidak boleh duplikasi pada table “Vihicle_data_Tab”
	var VehicleDataTab model.VehicleDataTab
	validasiEngineNo := r.db.Where("engine_no=?", vehicleEngineNo).First(&VehicleDataTab)

	return validasiEngineNo.Error
}

func ValidateChasisNo(vehicleChasisNo string) error {
	//	VEHICLE_CHASIS_NO Tidak Boleh Kosong
	if vehicleChasisNo == "" {
		return errors.New("vehicle chasis no kosong")
	}
	return nil
}

func (r *repository) ValidateDuplicateChasisNo(vehicleChasisNo string) error {
	//	VEHICLE_CHASIS_NO tidak boleh duplikasi pada table “Vihicle_data_Tab”
	var VehicleDataTab model.VehicleDataTab
	validasiChasisNo := r.db.Where("chasis_no=?", vehicleChasisNo).First(&VehicleDataTab)

	return validasiChasisNo.Error
}

func ConvertStringtoInt(input string) (int64, error) {
	if input == "" {
		return 0, nil
	}
	input = strings.TrimSpace(input)
	result, err := strconv.ParseInt(input, 10, 8)
	return result, err
}

func ConvertStringtoFloat(input string) (float64, error) {
	result, err := strconv.ParseFloat(input, 64)
	return result, err
}

func ConvertStringtoDate(input string) (time.Time, error) {
	result, err := time.Parse("2006-01-02", input)
	return result, err
}

func ConvertStringtoDateTime(input string) (time.Time, error) {
	result, err := time.Parse("2006-01-02 15:04:05", input)
	return result, err
}
