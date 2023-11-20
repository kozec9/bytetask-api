package services

import (
	"bytetask-api/internal/app/models"
	"bytetask-api/internal/app/repositories"
)

type TaskService interface {
	CreateTask(task models.Task) (*models.Task, error)
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

func (s *taskService) CreateTask(task models.Task) (*models.Task, error) {
	// Implement the logic to create a new task

	return s.TaskRepository.CreateTask(task)

}
