package model

import (
	"time"
)

// int8, int16, int32, int64
// float32, float64
type SkalaRentalTab struct {
	Custcode      string    `json:"custcode" gorm:"type:varchar(25); not null;"`
	Counter       int8      `json:"counter" gorm:"type:smallint"`
	OsBalance     string    `json:"osbalance" gorm:"type:money"`
	EndBalance    string    `json:"end_balance" gorm:"type:money"`
	DueDate       time.Time `json:"due_date" gorm:"type:timestamp"`
	EffRate       float64   `json:"eff_rate" gorm:"type:float"`
	Rental        string    `json:"rental" gorm:"type:money"`
	Principle     string    `json:"principle" gorm:"type:money"`
	Interest      string    `json:"interest" gorm:"type:money"`
	InputDate     time.Time `json:"inputdate" gorm:"type:timestamp"`
	InputBy       string    `json:"inputby" gorm:"type:varchar(50)"`
	LastModified  time.Time `json:"lastmodified" gorm:"type:timestamp"`
	ModifiedBy    string    `json:"modifiedby" gorm:"type:varchar(50)"`
	PaymentDate   time.Time `json:"payment_date" gorm:"type:timestamp"`
	Penalty       string    `json:"penalty" gorm:"type:money"`
	PaymentAmount string    `json:"payment_amount" gorm:"type:money"`
	PaymentType   int8      `json:"payment_type" gorm:"type:smallint"`
}

func (srt *SkalaRentalTab) TableName() string {
	return "skala_rental_tab"
}
