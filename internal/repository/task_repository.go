package repository

import (
	"context"
	"database/sql"
	"todo-list-service/internal/model"
)

type TaskRepository interface {
	GetAll(ctx context.Context, db *sql.DB) ([]model.Task, error)
	GetById(ctx context.Context, db *sql.DB, id int) (*model.Task, error)
	Create(ctx context.Context, tx *sql.Tx, task model.Task) (*int, error)
	Update(ctx context.Context, tx *sql.Tx, task model.Task) error
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}

type TaskRepositoryImpl struct {
}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (r *TaskRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, task model.Task) (*int, error) {
	// Define query
	query := `INSERT INTO tasks(title, description, due_date, is_active) VALUES ($1, $2, $3, 1) RETURNING id;`

	// Execute query
	row := tx.QueryRowContext(ctx, query, task.Title, task.Description, task.DueDate)

	// Get id
	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *TaskRepositoryImpl) GetAll(ctx context.Context, db *sql.DB) ([]model.Task, error) {
	// Define query
	query := `SELECT id, title, description, status, due_date, is_active, created_at, updated_at FROM tasks WHERE "is_active" = 1 ORDER BY due_date ASC;`

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	// Scan rows and store to slice
	var tasks []model.Task
	for rows.Next() {
		var task model.Task

		if err := rows.Scan(
			&task.Id, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.IsActive, &task.CreatedAt, &task.UpdatedAt,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepositoryImpl) GetById(ctx context.Context, db *sql.DB, id int) (*model.Task, error) {
	// Define query
	query := `SELECT id, title, description, status, due_date, is_active, created_at, updated_at FROM tasks WHERE id = $1;`

	// Execute query
	row := db.QueryRowContext(ctx, query, id)
	var task model.Task

	if err := row.Scan(
		&task.Id, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.IsActive, &task.CreatedAt, &task.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, task model.Task) error {
	// Define query
	query := `UPDATE tasks SET title=$1, description=$2, due_date=$3, status=$4, updated_at=CURRENT_TIMESTAMP WHERE id=$5;`

	// Execute query
	_, err := tx.ExecContext(ctx, query, task.Title, task.Description, task.DueDate, task.Status, task.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	// Define query
	query := `UPDATE tasks SET is_active = 0, updated_at = CURRENT_TIMESTAMP WHERE id = $1;`

	// Execute query
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
