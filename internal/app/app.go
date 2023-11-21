package app

import (
	"bytetask-api/configs"
	"bytetask-api/internal/app/handlers"
	"bytetask-api/internal/app/repositories"
	"bytetask-api/internal/app/services"

	"context"
	"fmt"
	"log"
	"net/http"

	taskroute "bytetask-api/internal/app/routes/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router *gin.Engine
	DB     *mongo.Database
	PORT   string
}

func NewApp() *App {
	cfg, err := configs.LoadConfig("../")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:" + cfg.PORT} // Replace with the origin of your frontend application
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	// Initialize MongoDB
	clientOptions := options.Client().ApplyURI(cfg.MONGOURL)
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

	db := client.Database(cfg.MONGONAME)

	collectionTask := db.Collection("task")

	// Initialize repositories, services, and handlers
	taskRepository := repositories.NewTaskRepository(collectionTask)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	app := &App{
		Router: r,
		DB:     db,
		PORT:   cfg.PORT,
	}

	fmt.Println(taskHandler)

	// Api Versioning v1
	routerV1 := r.Group("/api/v1")

	taskroute.SetupTaskRoutes(routerV1, taskHandler)

	app.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "WelCome to ApiByteFrost"})

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
