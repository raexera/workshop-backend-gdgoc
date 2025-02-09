package model

import "time"

type Task struct {
	Id          int
	Title       string
	Description *string
	Status      int
	DueDate     time.Time
	IsActive    int
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
