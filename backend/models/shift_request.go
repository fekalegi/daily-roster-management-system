package model

type ShiftRequest struct {
	ID       int    `json:"id"`
	WorkerID int    `json:"worker_id"`
	ShiftID  int    `json:"shift_id"`
	Status   string `json:"status"` // "pending", "approved", "rejected"
}

type ShiftRequestDetail struct {
	ID     int    `json:"id"`
	Status string `json:"status"`

	Worker *Worker `json:"worker"`
	Shift  *Shift  `json:"shift"`
}
