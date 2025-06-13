package repository

import (
	"task-service/internal/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAll(offset int, limit int, statusFilter string) ([]model.Task, error)
	GetByID(id uint) (*model.Task, error)
	Create(task *model.Task) error
	Update(task *model.Task) error
	Delete(task *model.Task) error
	Count(statusFilter string) (int64, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAll(offset int, limit int, statusFilter string) ([]model.Task, error) {
	var tasks []model.Task
	query := r.db.Model(&model.Task{})

	if statusFilter != "" {
		query = query.Where("status = ?", statusFilter)
	}

	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Count(statusFilter string) (int64, error) {
	var count int64
	query := r.db.Model(&model.Task{})

	if statusFilter != "" {
		query = query.Where("status = ?", statusFilter)
	}

	err := query.Count(&count).Error
	return count, err
}

func (r *taskRepository) GetByID(id uint) (*model.Task, error) {
	var task model.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(task *model.Task) error {
	return r.db.Delete(task).Error
}