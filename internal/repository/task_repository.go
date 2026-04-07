package repository

import (
	"errors"
	"go-repaso/internal/domain"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *domain.Task) error
	FindAllByUserID(userID uint) ([]domain.Task, error)
	FindByIDAndUserID(id uint, userID uint) (*domain.Task, error)
	FindByID(id uint) (*domain.Task, error)
	Update(task *domain.Task) error
	Delete(task *domain.Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) FindAllByUserID(userID uint) ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) FindByIDAndUserID(id uint, userID uint) (*domain.Task, error) {
	var task domain.Task
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrTaskNotFound
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) Update(task *domain.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(task *domain.Task) error {
	return r.db.Delete(task).Error
}

func (r *taskRepository) FindByID(id uint) (*domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrTaskNotFound
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}
