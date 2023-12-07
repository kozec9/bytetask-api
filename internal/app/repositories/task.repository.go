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
	UpdateTask(id string, task models.Task) (*models.TaskResponse, error)
	DeleteTask(id string) (*models.TaskResponse, error)
	GetTask(id string) (*models.TaskResponse, error)
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

// Update update a task in the database.
func (r *taskRepository) UpdateTask(id string, task models.Task) (*models.TaskResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("invalid ObjectID format")
	}

	// Update a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID}

	// Update the book into the MongoDB collection
	result, err := r.Collection.UpdateOne(ctx, filter, bson.M{"$set": task})
	if err != nil {
		return nil, err
	}

	// Check if any document was modified
	if result.ModifiedCount == 0 {
		return nil, errors.New("no document found for the given ID")
	}

	// Query the database to get the updated book object
	var updatedBookFromDB models.TaskResponse
	err = r.Collection.FindOne(ctx, filter).Decode(&updatedBookFromDB)
	if err != nil {
		return nil, err
	}

	return &updatedBookFromDB, nil
}

// Delete delete a task in the database.
func (r *taskRepository) DeleteTask(id string) (*models.TaskResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("invalid ObjectID format")
	}
	// Delete a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID}

	var deletedTask models.TaskResponse
	result := r.Collection.FindOneAndDelete(ctx, filter).Decode(&deletedTask)
	if result == mongo.ErrNoDocuments {
		return nil, errors.New("no document found for the given ID")
	} else if result != nil {
		return nil, result
	}

	return &deletedTask, nil
}

// Delete delete a task in the database.
func (r *taskRepository) GetTask(id string) (*models.TaskResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("invalid ObjectID format")
	}
	// Delete a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID}

	var task models.TaskResponse
	result := r.Collection.FindOne(ctx, filter).Decode(&task)
	if result == mongo.ErrNoDocuments {
		return nil, errors.New("no document found for the given ID")
	} else if result != nil {
		return nil, result
	}

	return &task, nil
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
