package model

import (
	"time"
)

// int8, int16, int32, int64
// float32, float64
type LoanDataTab struct {
	Custcode             string    `json:"custcode" gorm:"type:varchar(25); not null; unique; primaryKey"`
	Branch               string    `json:"branch" gorm:"type:varchar(50)"`
	OTR                  string    `json:"otr" gorm:"type:money"`
	DownPayment          string    `json:"down_payment" gorm:"type:money"`
	LoanAmount           string    `json:"loan_amount" gorm:"type:money"`
	LoanPeriod           string    `json:"loan_period" gorm:"type:varchar(6)"`
	InterestType         int8      `json:"interest_type" gorm:"type:smallint"`
	InterestFlat         float32   `json:"interest_flat" gorm:"type:real"`
	InterestEffective    float32   `json:"interest_effective" gorm:"type:real"`
	EffectivePaymentType int8      `json:"effective_payment_type" gorm:"type:smallint"`
	AdminFee             string    `json:"admin_fee" gorm:"type:money"`
	MonthlyPayment       string    `json:"monthly_payment" gorm:"type:money"`
	InputtDate           time.Time `json:"input_date" gorm:"type:timestamp"`
	LasttModified        time.Time `json:"last_modified" gorm:"type:timestamp"`
	ModifieddBy          string    `json:"modified_by" gorm:"type:varchar(20)"`
	InputDate            time.Time `json:"inputdate" gorm:"type:timestamp"`
	InputBy              string    `json:"inputby" gorm:"type:varchar(50)"`
	LastModified         time.Time `json:"lastmodified" gorm:"type:timestamp"`
	ModifiedBy           string    `json:"modifiedby" gorm:"type:varchar(50)"`
}

func (ldt *LoanDataTab) TableName() string {
	return "loan_data_tab"
}
