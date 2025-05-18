package usecase

import (
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/repository"
)

type WorkerUsecase interface {
	GetAll() ([]model.Worker, error)
}

type workerUsecase struct {
	// Could inject userRepo here to check DB for user info
	workerRepo repository.WorkerRepository
}

func NewWorkerUsecase(worker repository.WorkerRepository) WorkerUsecase {
	return &workerUsecase{
		workerRepo: worker,
	}
}

func (u *workerUsecase) GetAll() ([]model.Worker, error) {
	return u.workerRepo.GetAll()
}
