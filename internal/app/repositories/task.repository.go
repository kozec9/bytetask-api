package repositories

import (
	"bytetask-api/internal/app/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	CreateTask(task models.Task) (*models.Task, error)
	// Add other CRUD repository methods as needed
}

// TaskRepository implements the repository interface for the Task entity.
type taskRepository struct {
	Collection *mongo.Collection
}

// NewTaskRepository creates a new instance of TaskRepository.
func NewTaskRepository(collection *mongo.Collection) TaskRepository {
	return &taskRepository{
		Collection: collection,
	}
}

// Create creates a new task in the database.
func (r *taskRepository) CreateTask(task models.Task) (*models.Task, error) {
	// Implement the logic to create a new task in the database.
	// Use the r.db instance to perform the database operation.
	return &task, nil
}
