package repository

import (
	model "daily-worker-roster-management-system/models"
	"database/sql"
)

type shiftRequestRepository struct {
	DB *sql.DB
}

func NewShiftRequestRepository(db *sql.DB) ShiftRequestRepository {
	return &shiftRequestRepository{DB: db}
}

func (r *shiftRequestRepository) Create(req *model.ShiftRequest) error {
	_, err := r.DB.Exec(
		`INSERT INTO shift_requests (worker_id, shift_id, status) VALUES (?, ?, ?)`,
		req.WorkerID, req.ShiftID, req.Status,
	)
	return err
}

func (r *shiftRequestRepository) GetPending() ([]model.ShiftRequest, error) {
	rows, err := r.DB.Query("SELECT id, worker_id, shift_id, status FROM shift_requests WHERE status = 'pending'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []model.ShiftRequest
	for rows.Next() {
		var req model.ShiftRequest
		if err := rows.Scan(&req.ID, &req.WorkerID, &req.ShiftID, &req.Status); err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}

	return requests, nil
}

func (r *shiftRequestRepository) UpdateStatus(id int, status string) error {
	_, err := r.DB.Exec(`UPDATE shift_requests SET status = ? WHERE id = ?`, status, id)
	return err
}

func (r *shiftRequestRepository) GetByID(id int) (*model.ShiftRequest, error) {
	row := r.DB.QueryRow("SELECT id, worker_id, shift_id, status FROM shift_requests WHERE id = ?", id)

	var s model.ShiftRequest
	if err := row.Scan(&s.ID, &s.WorkerID, &s.ShiftID, &s.Status); err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *shiftRequestRepository) GetPendingWithDetails() ([]model.ShiftRequestDetail, error) {
	query := `
        SELECT 
            sr.id, sr.status,
            w.id, w.name,
            s.id, s.date, s.start_time, s.end_time, s.role, s.location
        FROM shift_requests sr
        JOIN workers w ON sr.worker_id = w.id
        JOIN shifts s ON sr.shift_id = s.id
        WHERE sr.status = 'pending'
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.ShiftRequestDetail
	for rows.Next() {
		detail := model.ShiftRequestDetail{
			Worker: &model.Worker{},
			Shift:  &model.Shift{},
		}
		if err := rows.Scan(
			&detail.ID, &detail.Status,
			&detail.Worker.ID, &detail.Worker.Name,
			&detail.Shift.ID, &detail.Shift.Date, &detail.Shift.StartTime, &detail.Shift.EndTime,
			&detail.Shift.Role, &detail.Shift.Location,
		); err != nil {
			return nil, err
		}
		results = append(results, detail)
	}

	return results, nil
}

func (r *shiftRequestRepository) GetWithDetails() ([]model.ShiftRequestDetail, error) {
	query := `
        SELECT 
            sr.id, sr.status,
            w.id, w.name,
            s.id, s.date, s.start_time, s.end_time, s.role, s.location
        FROM shift_requests sr
        JOIN workers w ON sr.worker_id = w.id
        JOIN shifts s ON sr.shift_id = s.id
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.ShiftRequestDetail
	for rows.Next() {
		detail := model.ShiftRequestDetail{
			Worker: &model.Worker{},
			Shift:  &model.Shift{},
		}
		if err := rows.Scan(
			&detail.ID, &detail.Status,
			&detail.Worker.ID, &detail.Worker.Name,
			&detail.Shift.ID, &detail.Shift.Date, &detail.Shift.StartTime, &detail.Shift.EndTime,
			&detail.Shift.Role, &detail.Shift.Location,
		); err != nil {
			return nil, err
		}
		results = append(results, detail)
	}

	return results, nil
}

func (r *shiftRequestRepository) GetWithDetailsFilterStatus(status string) ([]model.ShiftRequestDetail, error) {
	query := `
        SELECT 
            sr.id, sr.status,
            w.id, w.name,
            s.id, s.date, s.start_time, s.end_time, s.role, s.location
        FROM shift_requests sr
        JOIN workers w ON sr.worker_id = w.id
        JOIN shifts s ON sr.shift_id = s.id
        WHERE status = ?
    `
	rows, err := r.DB.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.ShiftRequestDetail
	for rows.Next() {
		detail := model.ShiftRequestDetail{
			Worker: &model.Worker{},
			Shift:  &model.Shift{},
		}
		if err := rows.Scan(
			&detail.ID, &detail.Status,
			&detail.Worker.ID, &detail.Worker.Name,
			&detail.Shift.ID, &detail.Shift.Date, &detail.Shift.StartTime, &detail.Shift.EndTime,
			&detail.Shift.Role, &detail.Shift.Location,
		); err != nil {
			return nil, err
		}
		results = append(results, detail)
	}

	return results, nil
}
