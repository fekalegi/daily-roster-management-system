package repository

import (
	model "daily-worker-roster-management-system/models"
	"database/sql"
)

type shiftRepository struct {
	DB *sql.DB
}

func NewShiftRepository(db *sql.DB) ShiftRepository {
	return &shiftRepository{DB: db}
}

func (r *shiftRepository) GetAll() ([]model.Shift, error) {
	rows, err := r.DB.Query("SELECT id, date, start_time, end_time, role, location FROM shifts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []model.Shift
	for rows.Next() {
		var s model.Shift
		if err := rows.Scan(&s.ID, &s.Date, &s.StartTime, &s.EndTime, &s.Role, &s.Location); err != nil {
			return nil, err
		}
		shifts = append(shifts, s)
	}

	return shifts, nil
}

func (r *shiftRepository) Create(s *model.Shift) error {
	_, err := r.DB.Exec(
		"INSERT INTO shifts (date, start_time, end_time, role, location) VALUES (?, ?, ?, ?, ?)",
		s.Date, s.StartTime, s.EndTime, s.Role, s.Location,
	)
	return err
}

func (r *shiftRepository) GetByID(id int) (*model.Shift, error) {
	row := r.DB.QueryRow("SELECT id, date, start_time, end_time, role, location FROM shifts WHERE id = ?", id)

	var s model.Shift
	if err := row.Scan(&s.ID, &s.Date, &s.StartTime, &s.EndTime, &s.Role, &s.Location); err != nil {
		return nil, err
	}
	return &s, nil
}
