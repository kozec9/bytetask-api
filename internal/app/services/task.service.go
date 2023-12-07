package services

import (
	"bytetask-api/internal/app/models"
	"bytetask-api/internal/app/repositories"
)

type TaskService interface {
	CreateTask(task models.Task) (*models.TaskResponse, error)
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

func (s *taskService) GetTasks() (*[]models.TaskResponse, error) {
	return s.TaskRepository.GetTasks()
}
