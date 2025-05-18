package model

type Shift struct {
    ID        int    `json:"id"`
    Date      string `json:"date"`
    StartTime string `json:"start_time"`
    EndTime   string `json:"end_time"`
    Role      string `json:"role"`
    Location  string `json:"location"`
}
