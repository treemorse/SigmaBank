package models

type Task struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	PointID     uint    `json:"point_id"`
	WorkerID    uint    `json:"worker_id"`
	WorkerGrade string  `json:"worker_grade"`
	Duration    float64 `json:"duration"` // in hours
	Priority    int     `json:"priority"`
}
