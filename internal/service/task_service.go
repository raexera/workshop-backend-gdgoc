package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo-list-service/internal/model"
	"todo-list-service/internal/model/request"
	"todo-list-service/internal/model/response"
	"todo-list-service/internal/repository"
	"todo-list-service/pkg"
)

type TaskService interface {
	GetAll(ctx context.Context) ([]response.TaskResponse, error)
	GetById(ctx context.Context, id int) (*response.TaskResponse, error)
	Create(ctx context.Context, request request.CreateTaskRequest) (*response.TaskResponse, error)
	Update(ctx context.Context, id int, request request.UpdateTaskRequest) error
	Delete(ctx context.Context, id int) error
	constructTaskResponse(task model.Task) response.TaskResponse
	constructCreateTask(request request.CreateTaskRequest) model.Task
	constructUpdateTask(id int, request request.UpdateTaskRequest) model.Task
}

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
	DB             *sql.DB
}

func NewTaskService(taskRepository repository.TaskRepository, db *sql.DB) TaskService {
	return &TaskServiceImpl{
		TaskRepository: taskRepository,
		DB:             db,
	}
}

func (s *TaskServiceImpl) Create(ctx context.Context, request request.CreateTaskRequest) (*response.TaskResponse, error) {
	// Begin Transaction
	tx, err := s.DB.Begin()
	if err != nil {
		fmt.Println("[internal][service][task][create] error begin transaction:", err)
		return nil, pkg.ErrInternalServerError
	}

	// Defer rollback
	defer func() error {
		if err != nil {
			tx.Rollback()
			return pkg.ErrInternalServerError
		}

		return nil
	}()

	// Create task
	task := s.constructCreateTask(request)
	id, err := s.TaskRepository.Create(ctx, tx, task)
	if err != nil {
		fmt.Println("[internal][service][task][create]:", err)
		return nil, pkg.ErrInternalServerError
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, pkg.ErrInternalServerError
	}

	// Get task by id
	newTask, err := s.GetById(ctx, *id)
	if err != nil {
		fmt.Println("[internal][service][task][create]:", err)
		return nil, pkg.ErrNotFound
	}

	return newTask, nil
}

func (s *TaskServiceImpl) GetAll(ctx context.Context) ([]response.TaskResponse, error) {
	tasks, err := s.TaskRepository.GetAll(ctx, s.DB)
	if err != nil {
		fmt.Println("[internal][service][task][get-all]:", err)
		return nil, pkg.ErrNotFound
	}

	var responses []response.TaskResponse
	for _, task := range tasks {
		response := s.constructTaskResponse(task)
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *TaskServiceImpl) GetById(ctx context.Context, id int) (*response.TaskResponse, error) {
	task, err := s.TaskRepository.GetById(ctx, s.DB, id)
	if err != nil {
		fmt.Println("[internal][service][task][get-by-id]:", err)
		return nil, pkg.ErrNotFound
	}

	response := s.constructTaskResponse(*task)
	return &response, nil
}

func (s *TaskServiceImpl) Update(ctx context.Context, id int, request request.UpdateTaskRequest) error {
	// Check task dengan id = n apakah ada?
	data, err := s.GetById(ctx, id)
	if err != nil || data == nil {
		return pkg.ErrNotFound
	}

	// Begin Transaction
	tx, err := s.DB.Begin()
	if err != nil {
		fmt.Println("[internal][service][task][update] error begin transaction:", err)
		return pkg.ErrInternalServerError
	}

	// Defer commit or rollback
	defer func() error {
		if err != nil {
			tx.Rollback()
			return pkg.ErrInternalServerError
		}

		if err := tx.Commit(); err != nil {
			tx.Rollback()
			return pkg.ErrInternalServerError
		}

		return nil
	}()

	// Update task
	task := s.constructUpdateTask(id, request)
	if err := s.TaskRepository.Update(ctx, tx, task); err != nil {
		fmt.Println("[internal][service][task][update]:", err)
		return pkg.ErrInternalServerError
	}

	return nil
}

func (s *TaskServiceImpl) Delete(ctx context.Context, id int) error {
	// Check task dengan id = n apakah ada?
	data, err := s.GetById(ctx, id)
	if err != nil || data == nil {
		return pkg.ErrNotFound
	}

	// Begin Transaction
	tx, err := s.DB.Begin()
	if err != nil {
		fmt.Println("[internal][service][task][delete] error begin transaction:", err)
		return pkg.ErrInternalServerError
	}

	// Defer commit or rollback
	defer func() error {
		if err != nil {
			tx.Rollback()
			return pkg.ErrInternalServerError
		}

		if err := tx.Commit(); err != nil {
			tx.Rollback()
			return pkg.ErrInternalServerError
		}

		return nil
	}()

	// Delete task
	if err := s.TaskRepository.Delete(ctx, tx, id); err != nil {
		fmt.Println("[internal][service][task][delete] error begin transaction:", err)
		return pkg.ErrInternalServerError
	}

	return nil
}

func (s *TaskServiceImpl) constructTaskResponse(task model.Task) response.TaskResponse {
	return response.TaskResponse{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func (s *TaskServiceImpl) constructCreateTask(request request.CreateTaskRequest) model.Task {
	return model.Task{
		Id:          0,
		Title:       request.Title,
		Description: request.Description,
		DueDate:     request.DueDate,
		Status:      0,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
}

func (s *TaskServiceImpl) constructUpdateTask(id int, request request.UpdateTaskRequest) model.Task {
	return model.Task{
		Id:          id,
		Title:       request.Title,
		Description: request.Description,
		DueDate:     request.DueDate,
		Status:      request.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
}
