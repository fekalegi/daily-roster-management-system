package usecase

import (
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/repository"
	"fmt"
)

type AssignmentUsecase interface {
	Assign(a *model.Assignment) error
	GetByWorker(workerID int) ([]model.AssignmentDetail, error)
	GetByDate(date string) ([]model.AssignmentDetail, error)
	FindAll() ([]model.AssignmentDetail, error)
}

type assignmentUsecase struct {
	repo      repository.AssignmentRepository
	shiftRepo repository.ShiftRepository
}

func NewAssignmentUsecase(
	repo repository.AssignmentRepository,
	shiftRepo repository.ShiftRepository,
) AssignmentUsecase {
	return &assignmentUsecase{
		repo:      repo,
		shiftRepo: shiftRepo,
	}
}

func (u *assignmentUsecase) Assign(a *model.Assignment) error {
	shift, err := u.shiftRepo.GetByID(a.ShiftID)
	if err != nil {
		return err
	}
	err = u.checkShiftRequest(a.WorkerID, shift)
	if err != nil {
		return err
	}
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

func (u *assignmentUsecase) checkShiftRequest(workerId int, shift *model.Shift) error {
	// Check if shift already assigned
	assigned, err := u.repo.IsShiftAlreadyAssigned(shift.ID)
	if err != nil {
		return err
	}
	if assigned {
		return fmt.Errorf("shift already assigned")
	}

	// Check if worker has shift on the same day
	hasShift, err := u.repo.HasAssignmentOnDay(workerId, shift.Date)
	if err != nil {
		return err
	}
	if hasShift {
		return fmt.Errorf("worker already has shift that day")
	}

	// Check weekly limit (5 shifts per week)
	weekStart, weekEnd := weekBounds(shift.Date)
	count, err := u.repo.CountAssignmentsInWeek(workerId, weekStart, weekEnd)
	if err != nil {
		return err
	}
	if count >= 5 {
		return fmt.Errorf("worker already has 5 shifts this week")
	}
	return nil
}
