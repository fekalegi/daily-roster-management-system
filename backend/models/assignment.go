package model

type Assignment struct {
	ID       int `json:"id"`
	ShiftID  int `json:"shift_id"`
	WorkerID int `json:"worker_id"`
}

type AssignmentDetail struct {
	ID     int     `json:"id"`
	Shift  *Shift  `json:"shift"`
	Worker *Worker `json:"worker"`
}
