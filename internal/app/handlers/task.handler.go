package handlers

import (
	"bytetask-api/internal/app/models"
	"bytetask-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TaskHandler represents the handler for tasks.
type TaskHandler struct {
	// Add any dependencies or data fields needed for the handler.
	TaskService services.TaskService
}

// NewTaskHandler creates a new instance of TaskHandler.
func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := h.TaskService.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, createdBook)
}
