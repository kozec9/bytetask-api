package services

import (
	"bytetask-api/internal/app/models"
	"bytetask-api/internal/app/repositories"
)

type TaskService interface {
	CreateTask(task models.Task) (*models.TaskResponse, error)
	UpdateTask(id string, task models.Task) (*models.TaskResponse, error)
	DeleteTask(id string) (*models.TaskResponse, error)
	GetTask(id string) (*models.TaskResponse, error)
	GetTasks() (*[]models.TaskResponse, error)
	// Add other CRUD service methods as needed
}

type taskService struct {
	TaskRepository repositories.TaskRepository
}

func NewTaskService(taskRepository repositories.TaskRepository) TaskService {
	return &taskService{
		TaskRepository: taskRepository,
	}
}

func (s *taskService) CreateTask(task models.Task) (*models.TaskResponse, error) {
	return s.TaskRepository.CreateTask(task)
}

func (s *taskService) UpdateTask(id string, task models.Task) (*models.TaskResponse, error) {
	return s.TaskRepository.UpdateTask(id, task)
}

// DeleteTask implements TaskService.
func (s *taskService) DeleteTask(id string) (*models.TaskResponse, error) {
	return s.TaskRepository.DeleteTask(id)
}

// GetTask implements TaskService.
func (s *taskService) GetTask(id string) (*models.TaskResponse, error) {
	return s.TaskRepository.GetTask(id)
}

func (s *taskService) GetTasks() (*[]models.TaskResponse, error) {
	return s.TaskRepository.GetTasks()
}
