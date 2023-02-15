package model

import "time"

type StagingCustomer struct {
	Id                              int64     `json:"id" gorm:"type:bigint; not null;"`
	ScReff                          string    `json:"sc_reff" gorm:"type:varchar(80)"`
	ScCreateDate                    time.Time `json:"sc_create_date" gorm:"type:timestamp"`
	ScBranchCode                    string    `json:"sc_branch_code" gorm:"type:varchar(80)"`
	ScCompany                       string    `json:"sc_company" gorm:"type:varchar(80)"`
	ScFlag                          string    `json:"sc_flag" gorm:"type:char"`
	CustomerPpk                     string    `json:"customer_ppk" gorm:"type:varchar(80)"`
	CustomerName                    string    `json:"customer_name" gorm:"type:varchar(80)"`
	CustomerAddress1                string    `json:"customer_address1" gorm:"type:varchar(80)"`
	CustomerAddress2                string    `json:"customer_address2" gorm:"type:varchar(80)"`
	CustomerCity                    string    `json:"customer_city" gorm:"type:varchar(80)"`
	CustomerZip                     string    `json:"customer_zip" gorm:"type:varchar(80)"`
	CustomerBirthPlace              string    `json:"customer_birth_place" gorm:"type:varchar(80)"`
	CustomerBirthDate               string    `json:"customer_birth_date" gorm:"type:varchar(80)"`
	CustomerIDType                  string    `json:"customer_id_type" gorm:"type:varchar(80)"`
	CustomerIDNumber                string    `json:"customer_id_number" gorm:"type:varchar(80)"`
	CustomerMobileNo                string    `json:"customer_mobile_no" gorm:"type:varchar(80)"`
	CustomerMotherMaidenName        string    `json:"customer_mother_maiden_name" gorm:"type:varchar(80)"`
	LoanOtr                         string    `json:"loan_otr" gorm:"type:varchar(80)"`
	LoanDownPayment                 string    `json:"loan_down_payment" gorm:"type:varchar(80)"`
	LoanLoanAmountChanneling        string    `json:"loan_loan_amount_channeling" gorm:"type:varchar(80)"`
	LoanLoanPeriodChanneling        string    `json:"loan_loan_period_channeling" gorm:"type:varchar(80)"`
	LoanInterestFlatChanneling      string    `json:"loan_interest_flat_channeling" gorm:"type:varchar(80)"`
	LoanInterestEffectiveChanneling string    `json:"loan_interest_effective_channeling" gorm:"type:varchar(80)"`
	LoanEffectivePaymentType        string    `json:"loan_effective_payment_type" gorm:"type:varchar(80)"`
	LoanMonthlyPaymentChanneling    string    `json:"loan_monthly_payment_channeling" gorm:"type:varchar(80)"`
	LoanTglPk                       string    `json:"loan_tgl_pk" gorm:"type:varchar(80)"`
	LoanTglPkChanneling             string    `json:"loan_tgl_pk_channeling" gorm:"type:varchar(80)"`
	CollateralTypeID                string    `json:"collateral_type_id" gorm:"type:varchar(80)"`
	VehicleJenisProduk              string    `json:"vehicle_jenis_produk" gorm:"type:varchar(80)"`
	VehicleBrand                    string    `json:"vehicle_brand" gorm:"type:varchar(80)"`
	VehicleType                     string    `json:"vehicle_type" gorm:"type:varchar(150)"`
	VehicleYear                     string    `json:"vehicle_year" gorm:"type:varchar(80)"`
	VehicleJenis                    string    `json:"vehicle_jenis" gorm:"type:varchar(80)"`
	VehicleStatus                   string    `json:"vehicle_status" gorm:"type:varchar(80)"`
	VehicleColor                    string    `json:"vehicle_color" gorm:"type:varchar(80)"`
	VehiclePoliceNo                 string    `json:"vehicle_police_no" gorm:"type:varchar(80)"`
	VehicleEngineNo                 string    `json:"vehicle_engine_no" gorm:"type:varchar(80)"`
	VehicleChasisNo                 string    `json:"vehicle_chasis_no" gorm:"type:varchar(80)"`
	VehicleBpkb                     string    `json:"vehicle_bpkb" gorm:"type:varchar(80)"`
	VehicleStnk                     string    `json:"vehicle_stnk" gorm:"type:varchar(80)"`
	VehicleDealer                   string    `json:"vehicle_dealer" gorm:"type:varchar(80)"`
	VehicleAddressDealer1           string    `json:"vehicle_address_dealer1" gorm:"type:varchar(80)"`
	VehicleAddressDealer2           string    `json:"vehicle_address_dealer2" gorm:"type:varchar(80)"`
	VehicleCityDealer               string    `json:"vehicle_city_dealer" gorm:"type:varchar(80)"`
	VehicleTglStnk                  string    `json:"vehicle_tgl_stnk" gorm:"type:varchar(80)"`
	VehicleTglBpkb                  string    `json:"vehicle_tgl_bpkb" gorm:"type:varchar(80)"`
	VehicleUtilizationPurpose       string    `json:"vehicle_utilization_purpose" gorm:"type:varchar(80)"`
	LoanDrawdownKolektibilitas      string    `json:"loan_drawdown_kolektibilitas" gorm:"type:varchar(2)"`
	VehicleDealerID                 string    `json:"vehicle_dealer_id" gorm:"type:varchar(10)"`
}

func (sc *StagingCustomer) TableName() string {
	return "staging_customer"
}
