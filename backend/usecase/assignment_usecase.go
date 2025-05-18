package usecase

import (
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/repository"
)

type AssignmentUsecase interface {
	Assign(a *model.Assignment) error
	GetByWorker(workerID int) ([]model.AssignmentDetail, error)
	GetByDate(date string) ([]model.AssignmentDetail, error)
	FindAll() ([]model.AssignmentDetail, error)
}

type assignmentUsecase struct {
	repo repository.AssignmentRepository
}

func NewAssignmentUsecase(repo repository.AssignmentRepository) AssignmentUsecase {
	return &assignmentUsecase{repo: repo}
}

func (u *assignmentUsecase) Assign(a *model.Assignment) error {
	return u.repo.Create(a)
}

func (u *assignmentUsecase) GetByWorker(workerID int) ([]model.AssignmentDetail, error) {
	return u.repo.FindByWorkerID(workerID)
}

func (u *assignmentUsecase) GetByDate(date string) ([]model.AssignmentDetail, error) {
	return u.repo.FindByDate(date)
}

func (u *assignmentUsecase) FindAll() ([]model.AssignmentDetail, error) {
	return u.repo.FindAll()
}
