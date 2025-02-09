package response

import "time"

type TaskResponse struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      int        `json:"status"`
	DueDate     time.Time  `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json;"updated_at"`
}
