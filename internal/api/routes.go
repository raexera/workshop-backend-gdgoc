package api

import (
	"database/sql"
	"net/http"
	"todo-list-service/internal/handler"
	"todo-list-service/internal/repository"
	"todo-list-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	TaskHandler handler.TaskHandler
}

func InitRoutes(db *sql.DB) *gin.Engine {
	return setupRoutes(*initHandler(db))
}

func initHandler(db *sql.DB) *Handlers {
	// Validator
	validator := validator.New()

	// Task
	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository, db)
	taskHandler := handler.NewTaskHandler(taskService, validator)

	return &Handlers{
		TaskHandler: taskHandler,
	}
}

func setupRoutes(handler Handlers) *gin.Engine {
	route := gin.Default()

	api := route.Group("/api/v1")
	api.GET("ping", index)

	// Task
	task := api.Group("/tasks")
	task.GET("", handler.TaskHandler.GetAll)
	task.GET("/:id", handler.TaskHandler.GetById)
	task.POST("", handler.TaskHandler.Create)
	task.PUT("/:id", handler.TaskHandler.Update)
	task.DELETE("/:id", handler.TaskHandler.Delete)

	return route
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "pong",
	})
}
