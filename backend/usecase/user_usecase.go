package usecase

import (
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/repository"
	"errors"
)

type UserUsecase interface {
	ValidateCredentials(username, password string) (*model.User, error)
}

type userUsecase struct {
	// Could inject userRepo here to check DB for user info
	workerRepo repository.WorkerRepository
}

func NewUserUsecase(worker repository.WorkerRepository) UserUsecase {
	return &userUsecase{
		workerRepo: worker,
	}
}

func (u *userUsecase) ValidateCredentials(username, password string) (*model.User, error) {
	if username == "admin" && password == "password" {
		return &model.User{
			RoleId:   1,
			UserId:   0,
			Username: "Admin",
		}, nil
	}
	// Real workers
	if password == "password" {
		worker, err := u.workerRepo.GetByName(username)
		if err == nil {
			return &model.User{
				RoleId:   2,
				UserId:   worker.ID,
				Username: worker.Name,
			}, nil
		}
	}
	return nil, errors.New("invalid credentials")
}
