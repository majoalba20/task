package handler

import (
	"go-repaso/internal/domain"
	"go-repaso/internal/dto"
	"go-repaso/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func getUserID(c *gin.Context) (uint, error) {
	value, exists := c.Get("userID")
	if !exists {
		return 0, domain.ErrUnauthorized
	}

	userID, ok := value.(uint)
	if !ok {
		return 0, domain.ErrUnauthorized
	}

	return userID, nil
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		handleError(c, err)
		return
	}

	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, domain.ErrInvalidPayload)
		return
	}

	task, err := h.service.CreateTask(userID, req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		handleError(c, err)
		return
	}

	tasks, err := h.service.GetTasks(userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		handleError(c, err)
		return
	}

	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handleError(c, domain.ErrInvalidTaskID)
		return
	}

	task, err := h.service.GetTaskByID(userID, uint(taskID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		handleError(c, err)
		return
	}

	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handleError(c, domain.ErrInvalidTaskID)
		return
	}

	var req dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, domain.ErrInvalidPayload)
		return
	}

	task, err := h.service.UpdateTask(userID, uint(taskID), req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		handleError(c, err)
		return
	}

	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handleError(c, domain.ErrInvalidTaskID)
		return
	}

	err = h.service.DeleteTask(userID, uint(taskID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

func (h *TaskHandler) ProcessTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		handleError(c, err)
		return
	}

	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handleError(c, domain.ErrInvalidTaskID)
		return
	}

	task, err := h.service.ProcessTask(userID, uint(taskID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "task queued successfully",
		"task":    task,
	})
}

func handleError(c *gin.Context, err error) {
	switch err {
	case domain.ErrTaskNotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case domain.ErrInvalidStatus, domain.ErrInvalidPayload, domain.ErrInvalidTaskID:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case domain.ErrUnauthorized:
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case domain.ErrTaskCannotProcess:
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case domain.ErrTaskQueueIsFull:
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
