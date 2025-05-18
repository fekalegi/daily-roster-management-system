package repository

import model "daily-worker-roster-management-system/models"

//go:generate mockgen -destination=../mocks/repository/mock_repository.go -package=mock_repository -source=repository_interface.go

type AssignmentRepository interface {
	Create(a *model.Assignment) error
	FindByWorkerID(workerID int) ([]model.AssignmentDetail, error)
	FindByDate(date string) ([]model.AssignmentDetail, error)
	CountAssignmentsInWeek(workerID int, weekStart string, weekEnd string) (int, error)
	HasAssignmentOnDay(workerID int, date string) (bool, error)
	IsShiftAlreadyAssigned(shiftID int) (bool, error)
	FindAll() ([]model.AssignmentDetail, error)
}

type WorkerRepository interface {
	GetByName(name string) (*model.Worker, error)
	GetAll() ([]model.Worker, error)
}

type ShiftRepository interface {
	GetAll() ([]model.Shift, error)
	Create(shift *model.Shift) error
	GetByID(id int) (*model.Shift, error)
}

type ShiftRequestRepository interface {
	Create(req *model.ShiftRequest) error
	GetPending() ([]model.ShiftRequest, error)
	UpdateStatus(id int, status string) error
	GetByID(id int) (*model.ShiftRequest, error)
	GetPendingWithDetails() ([]model.ShiftRequestDetail, error)
	GetPastWithDetails() ([]model.ShiftRequestDetail, error)
}
