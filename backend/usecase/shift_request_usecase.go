package usecase

import (
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/repository"
	"errors"
	"fmt"
	"time"
)

type ShiftRequestUsecase interface {
	CreateRequest(req *model.ShiftRequest) error
	GetPendingRequests() ([]model.ShiftRequestDetail, error)
	Approve(id int) error
	Reject(id int) error
}

type shiftRequestUsecase struct {
	repo           repository.ShiftRequestRepository
	shiftRepo      repository.ShiftRepository
	assignmentRepo repository.AssignmentRepository
}

func NewShiftRequestUsecase(r repository.ShiftRequestRepository, shiftRepo repository.ShiftRepository, aRepo repository.AssignmentRepository) ShiftRequestUsecase {
	return &shiftRequestUsecase{repo: r, shiftRepo: shiftRepo, assignmentRepo: aRepo}
}

func (u *shiftRequestUsecase) CreateRequest(req *model.ShiftRequest) error {
	shift, err := u.shiftRepo.GetByID(req.ShiftID)
	if err != nil {
		return err
	}

	err = u.checkShiftRequest(req, shift)
	if err != nil {
		return err
	}

	return u.repo.Create(req)
}

func (u *shiftRequestUsecase) GetPendingRequests() ([]model.ShiftRequestDetail, error) {
	return u.repo.GetPendingWithDetails()
}

func (u *shiftRequestUsecase) Approve(id int) error {
	req, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	shift, err := u.shiftRepo.GetByID(req.ShiftID)
	if err != nil {
		return err
	}

	// Same conflict checks as above
	err = u.checkShiftRequest(req, shift)
	if err != nil {
		return err
	}

	err = u.assignmentRepo.Create(&model.Assignment{
		ShiftID:  req.ShiftID,
		WorkerID: req.WorkerID,
	})
	if err != nil {
		return err
	}

	return u.repo.UpdateStatus(req.ID, "approved")
}

func (u *shiftRequestUsecase) Reject(id int) error {
	req, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	if req.Status != "pending" {
		return errors.New("shift request is not pending")
	}

	return u.repo.UpdateStatus(id, "rejected")
}

func weekBounds(dateStr string) (string, string) {
	t, _ := time.Parse("2006-01-02", dateStr)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	} // Sunday → 7

	start := t.AddDate(0, 0, -weekday+1) // Monday
	end := start.AddDate(0, 0, 6)        // Sunday

	return start.Format("2006-01-02"), end.Format("2006-01-02")
}

func (u *shiftRequestUsecase) checkShiftRequest(req *model.ShiftRequest, shift *model.Shift) error {
	// Check if shift already assigned
	assigned, err := u.assignmentRepo.IsShiftAlreadyAssigned(req.ShiftID)
	if err != nil {
		return err
	}
	if assigned {
		return fmt.Errorf("shift already assigned")
	}

	// Check if worker has shift on the same day
	hasShift, err := u.assignmentRepo.HasAssignmentOnDay(req.WorkerID, shift.Date)
	if err != nil {
		return err
	}
	if hasShift {
		return fmt.Errorf("worker already has shift that day")
	}

	// Check weekly limit (5 shifts per week)
	weekStart, weekEnd := weekBounds(shift.Date)
	count, err := u.assignmentRepo.CountAssignmentsInWeek(req.WorkerID, weekStart, weekEnd)
	if err != nil {
		return err
	}
	if count >= 5 {
		return fmt.Errorf("worker already has 5 shifts this week")
	}
	return nil
}
