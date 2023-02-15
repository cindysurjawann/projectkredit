package model

type GeneralParameter struct {
	Id        int64  `json:"id" gorm:"type:bigint; not null; unique; auto_increment"`
	Parameter string `json:"parameter" gorm:"type:varchar(100); not null;"`
	Value     string `json:"value" gorm:"type:varchar(200);"`
}

func (gp *GeneralParameter) TableName() string {
	return "general_parameter"
}
