package model

import (
	"time"
)

// int8, int16, int32, int64
type CustomerDataTab struct {
	Custcode          string    `json:"custcode" gorm:"type:varchar(25); not null; unique"`
	PPK               string    `json:"ppk" gorm:"type:varchar(20)"`
	Name              string    `json:"name" gorm:"type:varchar(100)"`
	Address1          string    `json:"address1" gorm:"type:varchar(40)"`
	Address2          string    `json:"address2" gorm:"type:varchar(40)"`
	City              string    `json:"city" gorm:"type:varchar(30)"`
	Zip               string    `json:"zip" gorm:"type:varchar(6)"`
	BirthPlace        string    `json:"birth_place" gorm:"type:varchar(20)"`
	BirthDate         time.Time `json:"birth_date" gorm:"type:date"`
	IdType            int8      `json:"id_type" gorm:"type:smallint"`
	IdNumber          string    `json:"id_number" gorm:"type:varchar(30)"`
	MobileNo          string    `json:"mobile_no" gorm:"type:varchar(20)"`
	DrawdownDate      time.Time `json:"drawdown_date" gorm:"type:date"`
	TglPkChanneling   time.Time `json:"tgl_pk_channeling" gorm:"type:date"`
	MotherMaidenName  string    `json:"mother_maiden_name" gorm:"type:varchar(100)"`
	ChannelingCompany string    `json:"channeling_company" gorm:"type:varchar(100)"`
	ApprovalStatus    string    `json:"approval_status" gorm:"type:varchar(2)"`
}

func (cdt *CustomerDataTab) TableName() string {
	return "customer_data_tab"
}
