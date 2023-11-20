package app

import (
	"bytetask-api/internal/app/handlers"
	"bytetask-api/internal/app/repositories"
	"bytetask-api/internal/app/services"

	"context"
	"fmt"
	"log"
	"net/http"

	taskroute "bytetask-api/internal/app/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

type App struct {
	Router *gin.Engine
	DB     *mongo.Database
}

func NewApp() *App {
	r := gin.Default()

	// Initialize MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("bytetask")

	collectionTask := db.Collection("task")

	// Initialize repositories, services, and handlers
	taskRepository := repositories.NewTaskRepository(collectionTask)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	app := &App{
		Router: r,
		DB:     db,
	}

	fmt.Println(taskHandler)

	taskroute.SetupTaskRoutes(app.Router, taskHandler)

	app.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "WelCome to ApiByteFrost"})
		return
	})

	return app
}

func (app *App) Run(port string) {
	err := app.Router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) Close() {
	// Close resources like database connections when the application is shut down
	if err := app.DB.Client().Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}
