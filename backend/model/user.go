package model

type User struct {
	UserId   string `json:"user_id" gorm:"type:varchar(20); not null; unique;"`
	Password string `json:"password" gorm:"type:varchar(100); not null;"`
	Name     string `json:"name" gorm:"type:varchar(200); not null"`
	Email    string `json:"email" gorm:"type:varchar(200); not null"`
}

func (u *User) TableName() string {
	return "kredit_user"
}
