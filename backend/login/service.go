package login

import (
	"errors"
	"kredit/backend/model"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(data LoginRequest) (model.User, int, error)
	Register(data RegisterRequest) (model.User, int, error)
	FindUser(user_id string) (model.User, int, error)
}

type service struct {
	repo LoginRepository
}

func NewService(repo LoginRepository) *service {
	return &service{repo}
}

func (s *service) FindUser(user_id string) (model.User, int, error) {
	User, err := s.repo.FindUser(user_id)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, errors.New("error find user")
	}
	return User, http.StatusOK, nil
}

func (s *service) Login(data LoginRequest) (model.User, int, error) {
	userId := data.UserId
	User, err := s.repo.FindUser(userId)

	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "user not found" {
			status = http.StatusUnauthorized
			err = errors.New("user not found")
		}
		return model.User{}, status, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(data.Password))
	if err != nil {
		return model.User{}, http.StatusUnauthorized, errors.New("password is wrong")
	}
	User.Password = ""
	return User, http.StatusOK, nil
}

func (s *service) Register(data RegisterRequest) (model.User, int, error) {
	row, err := s.repo.FindUser(data.UserId)
	if err == nil && row.Name != "" {
		return model.User{}, http.StatusBadRequest, errors.New("duplicate data")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	User := model.User{
		UserId:   data.UserId,
		Password: string(passwordHash),
		Name:     data.Name,
		Email:    data.Email,
	}
	res, err := s.repo.AddUser(User)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}
