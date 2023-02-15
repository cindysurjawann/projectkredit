package login

import (
	"errors"
	"kredit/backend/model"

	"gorm.io/gorm"
)

type LoginRepository interface {
	FindUser(userId string) (model.User, error)
	AddUser(user model.User) (model.User, error)
	UpdatePassword(user model.User) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUser(userId string) (model.User, error) {
	var User model.User
	rows := r.db.Where("user_id=?", userId).Find(&User)
	if rows.Error != nil {
		if errors.Is(rows.Error, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, rows.Error
	}
	return User, nil
}

func (r *repository) AddUser(user model.User) (model.User, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}

	return user, nil
}

func (r *repository) UpdatePassword(user model.User) (model.User, error) {
	var User model.User
	res := r.db.Model(&User).Where("user_id=?", user.UserId).Update("password", user.Password)
	if res.Error != nil {
		return model.User{}, res.Error
	}

	return user, nil
}
