package usecase

import "errors"

type UserUsecase interface {
	ValidateCredentials(username, password string) (int, error)
}

type userUsecase struct {
	// Could inject userRepo here to check DB for user info
}

func NewUserUsecase() UserUsecase {
	return &userUsecase{}
}

func (u *userUsecase) ValidateCredentials(username, password string) (int, error) {
	if username == "admin" && password == "password" {
		return 1, nil
	}
	if username == "role" && password == "password" {
		return 2, nil
	}
	return 0, errors.New("invalid credentials")
}
