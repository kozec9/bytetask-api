package repositories

import (
	"bytetask-api/internal/app/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type TaskRepository interface {
	CreateTask(task models.Task) (*models.TaskResponse, error)
	GetTasks() (*[]models.TaskResponse, error)
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
func (r *taskRepository) CreateTask(task models.Task) (*models.TaskResponse, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert the book into the MongoDB collection
	result, err := r.Collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	// Get the inserted book's ID from the result
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted ID")
	}

	// Query the database to get the complete book object with the generated ID
	var createdTask models.TaskResponse
	err = r.Collection.FindOne(ctx, bson.M{"_id": insertedID}).Decode(&createdTask)
	if err != nil {
		return nil, err
	}

	return &createdTask, nil
}

func (r *taskRepository) GetTasks() (*[]models.TaskResponse, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find all tasks in the MongoDB collection
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode the results into a slice of tasks
	var tasks []models.TaskResponse
	for cursor.Next(ctx) {
		var task models.TaskResponse
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return &tasks, nil
}
