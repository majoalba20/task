package dto

import "go-repaso/internal/domain"

type CreateTaskRequest struct {
	Title        string            `json:"title" binding:"required,min=3,max=255"`
	Description  string            `json:"description"`
	Status       domain.TaskStatus `json:"status"`
	Result       *string           `json:"result"`
	ErrorMessage *string           `json:"error_message"`
}

type UpdateTaskRequest struct {
	Title        *string            `json:"title"`
	Description  *string            `json:"description"`
	Status       *domain.TaskStatus `json:"status"`
	Result       *string            `json:"result"`
	ErrorMessage *string            `json:"error_message"`
}
