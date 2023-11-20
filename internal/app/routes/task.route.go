package routes

import (
	"bytetask-api/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

// initTaskRoutes initializes the task routes and registers the necessary handlers
func SetupTaskRoutes(router *gin.Engine, taskHandler *handlers.TaskHandler) {
	taskGroup := router.Group("/tasks")

	taskGroup.POST("", taskHandler.CreateTask)

}
