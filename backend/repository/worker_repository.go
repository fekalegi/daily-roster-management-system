package repository

import (
	model "daily-worker-roster-management-system/models"
	"database/sql"
)

type workerRepository struct {
	DB *sql.DB
}

type WorkerRepository interface {
	GetByName(name string) (*model.Worker, error)
}

func NewWorkerRequestRepository(db *sql.DB) WorkerRepository {
	return &workerRepository{DB: db}
}

func (r *workerRepository) GetByName(name string) (*model.Worker, error) {
	row := r.DB.QueryRow("SELECT id, name FROM workers WHERE name = ?", name)

	var worker model.Worker
	if err := row.Scan(&worker.ID, &worker.Name); err != nil {
		return nil, err
	}
	return &worker, nil
}
