package repository

import (
	model "daily-worker-roster-management-system/models"
	"database/sql"
)

type workerRepository struct {
	DB *sql.DB
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

func (r *workerRepository) GetAll() ([]model.Worker, error) {
	rows, err := r.DB.Query("SELECT id, name FROM workers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workers []model.Worker
	for rows.Next() {
		var s model.Worker
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		workers = append(workers, s)
	}

	return workers, nil
}
