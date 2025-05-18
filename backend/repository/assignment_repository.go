package repository

import (
	model "daily-worker-roster-management-system/models"
	"database/sql"
)

type assignmentRepository struct {
	DB *sql.DB
}

func NewAssignmentRepository(db *sql.DB) AssignmentRepository {
	return &assignmentRepository{DB: db}
}

func (r *assignmentRepository) Create(a *model.Assignment) error {
	_, err := r.DB.Exec(
		"INSERT INTO assignments (shift_id, worker_id) VALUES (?, ?)",
		a.ShiftID, a.WorkerID,
	)
	return err
}

func (r *assignmentRepository) FindByWorkerID(workerID int) ([]model.AssignmentDetail, error) {
	query := `
        SELECT 
            a.id,
            w.id, w.name,
            s.id, s.date, s.start_time, s.end_time, s.role, s.location
        FROM assignments a
        JOIN workers w ON a.worker_id = w.id
        JOIN shifts s ON a.shift_id = s.id 
		WHERE worker_id = ?`
	rows, err := r.DB.Query(query, workerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.AssignmentDetail
	for rows.Next() {
		detail := model.AssignmentDetail{
			Worker: &model.Worker{},
			Shift:  &model.Shift{},
		}
		if err := rows.Scan(
			&detail.ID,
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

func (r *assignmentRepository) FindByDate(date string) ([]model.AssignmentDetail, error) {
	query := `
        SELECT 
            a.id,
            w.id, w.name,
            s.id, s.date, s.start_time, s.end_time, s.role, s.location
        FROM assignments a
        JOIN workers w ON a.worker_id = w.id
        JOIN shifts s ON a.shift_id = s.id
        WHERE s.date = ?
    `
	rows, err := r.DB.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.AssignmentDetail
	for rows.Next() {
		detail := model.AssignmentDetail{
			Worker: &model.Worker{},
			Shift:  &model.Shift{},
		}
		if err := rows.Scan(
			&detail.ID,
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

func (r *assignmentRepository) CountAssignmentsInWeek(workerID int, weekStart string, weekEnd string) (int, error) {
	row := r.DB.QueryRow(`
        SELECT COUNT(*) FROM assignments a 
        JOIN shifts s ON a.shift_id = s.id 
        WHERE a.worker_id = ? AND s.date BETWEEN ? AND ?
    `, workerID, weekStart, weekEnd)

	var count int
	err := row.Scan(&count)
	return count, err
}

func (r *assignmentRepository) HasAssignmentOnDay(workerID int, date string) (bool, error) {
	row := r.DB.QueryRow(`
        SELECT COUNT(*) FROM assignments a 
        JOIN shifts s ON a.shift_id = s.id 
        WHERE a.worker_id = ? AND s.date = ?
    `, workerID, date)

	var count int
	err := row.Scan(&count)
	return count > 0, err
}

func (r *assignmentRepository) IsShiftAlreadyAssigned(shiftID int) (bool, error) {
	row := r.DB.QueryRow("SELECT COUNT(*) FROM assignments WHERE shift_id = ?", shiftID)
	var count int
	err := row.Scan(&count)
	return count > 0, err
}

func (r *assignmentRepository) FindAll() ([]model.AssignmentDetail, error) {
	query := `
        SELECT 
            a.id,
            w.id, w.name,
            s.id, s.date, s.start_time, s.end_time, s.role, s.location
        FROM assignments a
        JOIN workers w ON a.worker_id = w.id
        JOIN shifts s ON a.shift_id = s.id
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.AssignmentDetail
	for rows.Next() {
		detail := model.AssignmentDetail{
			Worker: &model.Worker{},
			Shift:  &model.Shift{},
		}
		if err := rows.Scan(
			&detail.ID,
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
