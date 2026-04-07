package service

import (
	"go-repaso/internal/domain"
	"go-repaso/internal/dto"
	"go-repaso/internal/queue"
	"go-repaso/internal/repository"
)

type TaskService interface {
	CreateTask(userID uint, req dto.CreateTaskRequest) (*domain.Task, error)
	GetTasks(userID uint) ([]domain.Task, error)
	GetTaskByID(userID, taskID uint) (*domain.Task, error)
	UpdateTask(userID, taskID uint, req dto.UpdateTaskRequest) (*domain.Task, error)
	DeleteTask(userID, taskID uint) error
	ProcessTask(userID, taskID uint) (*domain.Task, error)
}

type taskService struct {
	repo  repository.TaskRepository
	queue queue.TaskQueue
}

func NewTaskService(repo repository.TaskRepository, q queue.TaskQueue) TaskService {
	return &taskService{
		repo:  repo,
		queue: q,
	}
}

func isValidStatus(status domain.TaskStatus) bool {
	switch status {
	case domain.StatusPending,
		domain.StatusQueued,
		domain.StatusProcessing,
		domain.StatusDone,
		domain.StatusFailed:
		return true
	default:
		return false
	}
}

func (s *taskService) CreateTask(userID uint, req dto.CreateTaskRequest) (*domain.Task, error) {
	status := req.Status
	if status == "" {
		status = domain.StatusPending
	}

	if !isValidStatus(status) {
		return nil, domain.ErrInvalidStatus
	}

	task := &domain.Task{
		UserID:       userID,
		Title:        req.Title,
		Description:  req.Description,
		Status:       status,
		Result:       req.Result,
		ErrorMessage: req.ErrorMessage,
	}

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetTasks(userID uint) ([]domain.Task, error) {
	return s.repo.FindAllByUserID(userID)
}

func (s *taskService) GetTaskByID(userID, taskID uint) (*domain.Task, error) {
	return s.repo.FindByIDAndUserID(taskID, userID)
}

func (s *taskService) UpdateTask(userID, taskID uint, req dto.UpdateTaskRequest) (*domain.Task, error) {
	task, err := s.repo.FindByIDAndUserID(taskID, userID)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Status != nil {
		if !isValidStatus(*req.Status) {
			return nil, domain.ErrInvalidStatus
		}
		task.Status = *req.Status
	}
	if req.Result != nil {
		task.Result = req.Result
	}
	if req.ErrorMessage != nil {
		task.ErrorMessage = req.ErrorMessage
	}

	if err := s.repo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(userID, taskID uint) error {
	task, err := s.repo.FindByIDAndUserID(taskID, userID)
	if err != nil {
		return err
	}

	return s.repo.Delete(task)
}

func (s *taskService) ProcessTask(userID, taskID uint) (*domain.Task, error) {
	task, err := s.repo.FindByIDAndUserID(taskID, userID)
	if err != nil {
		return nil, err
	}

	if task.Status != domain.StatusPending && task.Status != domain.StatusFailed {
		return nil, domain.ErrTaskCannotProcess
	}

	task.Status = domain.StatusQueued
	task.Result = nil
	task.ErrorMessage = nil

	if err := s.repo.Update(task); err != nil {
		return nil, err
	}

	if err := s.queue.Enqueue(task.ID); err != nil {
		task.Status = domain.StatusFailed
		msg := "failed to enqueue task"
		task.ErrorMessage = &msg
		_ = s.repo.Update(task)

		return nil, domain.ErrTaskQueueIsFull
	}

	return task, nil
}
