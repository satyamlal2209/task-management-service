package service

import (
	"errors"
	"task-service/internal/model"
	"task-service/internal/repository"
)

type TaskService interface {
	ListTasks(page int, size int, statusFilter string) ([]model.Task, int64, error)
	GetTask(id uint) (*model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(id uint, updatedTask *model.Task) error
	DeleteTask(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) ListTasks(page int, size int, statusFilter string) ([]model.Task, int64, error) {
	offset := (page - 1) * size
	tasks, err := s.repo.GetAll(offset, size, statusFilter)
	if err != nil {
		return nil, 0, err
	}

	totalCount, err := s.repo.Count(statusFilter)
	if err != nil {
		return nil, 0, err
	}

	return tasks, totalCount, nil
}

func (s *taskService) GetTask(id uint) (*model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) CreateTask(task *model.Task) error {
	return s.repo.Create(task)
}

func (s *taskService) UpdateTask(id uint, updatedTask *model.Task) error {
	existingTask, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	existingTask.Title = updatedTask.Title
	existingTask.Description = updatedTask.Description
	existingTask.Status = updatedTask.Status
	existingTask.DueDate = updatedTask.DueDate

	return s.repo.Update(existingTask)
}

func (s *taskService) DeleteTask(id uint) error {
	existingTask, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if existingTask == nil {
		return errors.New("task not found")
	}

	return s.repo.Delete(existingTask)
}
