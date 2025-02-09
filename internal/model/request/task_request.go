package request

import "time"

type CreateTaskRequest struct {
	Title       string    `json:"title" validate:"required,max=255"`
	Description *string   `json:"description"`
	DueDate     time.Time `json:"due_date" validate:"required"`
}

type UpdateTaskRequest struct {
	Title       string    `json:"title" validate:"required,max=255"`
	Description *string   `json:"description"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	Status      int       `json:"status" validate:"required"`
}
