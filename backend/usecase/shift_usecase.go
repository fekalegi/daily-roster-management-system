package usecase

import (
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/repository"
)

type ShiftUsecase interface {
	GetShifts() ([]model.Shift, error)
	CreateShift(shift *model.Shift) error
}

type shiftUsecase struct {
	repo repository.ShiftRepository
}

func NewShiftUsecase(repo repository.ShiftRepository) ShiftUsecase {
	return &shiftUsecase{repo: repo}
}

func (u *shiftUsecase) GetShifts() ([]model.Shift, error) {
	return u.repo.GetAll()
}

func (u *shiftUsecase) CreateShift(shift *model.Shift) error {
	return u.repo.Create(shift)
}
