package interfaces

import "time"

type Status string

const (
	StatusStarted    Status = "started"
	StatusInProgress Status = "in-progress"
	StatusEnded      Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"Description"`
	Status      Status    `json:"Status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}
