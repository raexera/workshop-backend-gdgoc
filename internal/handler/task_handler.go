package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todo-list-service/internal/model/request"
	"todo-list-service/internal/service"
	"todo-list-service/pkg"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TaskHandler interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type TaskHandlerImpl struct {
	TaskService service.TaskService
	Validator   *validator.Validate
}

func NewTaskHandler(taskService service.TaskService, validator *validator.Validate) TaskHandler {
	return &TaskHandlerImpl{
		TaskService: taskService,
		Validator:   validator,
	}
}

func (h *TaskHandlerImpl) Create(c *gin.Context) {
	var request request.CreateTaskRequest
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("[internal][handler][task][create]:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": pkg.ErrBadRequest.Error(),
		})
		return
	}

	if err := h.Validator.Struct(request); err != nil {
		fmt.Println("[internal][handler][task][create]:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": pkg.ErrBadRequest.Error(),
		})
		return
	}

	data, err := h.TaskService.Create(ctx, request)
	if err != nil {
		fmt.Println("[internal][handler][task][create]:", err)
		switch {
		case errors.Is(err, pkg.ErrInternalServerError):
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "failed to create new task",
			})
		case errors.Is(err, pkg.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": pkg.ErrNotFound.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success create new task",
			"data":    data,
		})
	}
}

func (h *TaskHandlerImpl) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := h.TaskService.GetAll(ctx)
	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrInternalServerError):
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "failed to get all tasks",
			})
		case errors.Is(err, pkg.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": pkg.ErrNotFound.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success get all data",
			"data":    data,
		})
	}
}

func (h *TaskHandlerImpl) GetById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": pkg.ErrBadRequest.Error(),
		})
		return
	}

	data, err := h.TaskService.GetById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrInternalServerError):
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": pkg.ErrInternalServerError.Error(),
			})
		case errors.Is(err, pkg.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": pkg.ErrNotFound.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success get data by id",
			"data":    data,
		})
	}
}

func (h *TaskHandlerImpl) Update(c *gin.Context) {
	var request request.UpdateTaskRequest
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": pkg.ErrBadRequest.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": pkg.ErrBadRequest.Error(),
		})
		return
	}

	if err := h.Validator.Struct(request); err != nil {
		fmt.Println("[internal][handler][task][update]:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": pkg.ErrBadRequest.Error(),
		})
		return
	}

	if err := h.TaskService.Update(ctx, id, request); err != nil {
		fmt.Println("[internal][handler][task][update]:", err)
		switch {
		case errors.Is(err, pkg.ErrInternalServerError):
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "failed to update new task",
			})
		case errors.Is(err, pkg.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": pkg.ErrNotFound.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success update new task",
		})
	}
}

func (h *TaskHandlerImpl) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": pkg.ErrBadRequest.Error(),
		})
		return
	}

	if err := h.TaskService.Delete(ctx, id); err != nil {
		fmt.Println("[internal][handler][task][delete]:", err)
		switch {
		case errors.Is(err, pkg.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": pkg.ErrNotFound.Error(),
			})
		case errors.Is(err, pkg.ErrInternalServerError):
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": pkg.ErrInternalServerError.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success delete task",
		})
	}
}
