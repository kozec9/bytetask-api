package app

import (
	"bytetask-api/configs"
	"bytetask-api/internal/app/models"
	"net/http"
	"sync"
	"time"

	"context"
	"fmt"
	"log"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type App struct {
	Router *gin.Engine
	DB     *mongo.Database
	PORT   string
}

var (
	memoryMap = make(map[string]interface{})
	mutex     sync.RWMutex
)

func connectMongo(mongoURL string, mongoName string) (*mongo.Database, error) {

	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDBsssssss!")
	fmt.Println(mongoName)

	db := client.Database(mongoName)

	return db, nil
}

func NewApp() *App {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	r := gin.Default()

	fmt.Println("PORT: ", cfg.PORT)

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:" + cfg.PORT} // Replace with the origin of your frontend application
	// config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	// r.Use(cors.New(config))

	r.Use(corsMiddleware())

	// // Initialize MongoDB
	// clientOptions := options.Client().ApplyURI(cfg.MONGOURL)
	// client, err := mongo.Connect(context.Background(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Check the connection
	// err = client.Ping(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to MongoDB!")

	// db := client.Database(cfg.MONGONAME)

	// collectionTask := db.Collection("task")

	// // Initialize repositories, services, and handlers
	// taskRepository := repositories.NewTaskRepository(collectionTask)
	// taskService := services.NewTaskService(taskRepository)
	// taskHandler := handlers.NewTaskHandler(taskService)

	app := &App{
		Router: r,
		// DB:     db,
		PORT: cfg.PORT,
	}

	// fmt.Println(taskHandler)

	// Api Versioning v1
	routerV1 := r.Group("/api/v1")

	// taskroute.SetupTaskRoutes(routerV1, taskHandler)

	routerV1.GET("/config", func(c *gin.Context) {

		mutex.RLock()
		url, existsa := memoryMap["mongodbURL"]
		nmae, existsb := memoryMap["mongodbName"]
		mutex.RUnlock()

		if !existsa {
			setting := models.Setting{
				MongoURL:     "",
				DatabaseName: "",
			}

			c.JSON(http.StatusOK, setting)
			return
		}

		if !existsb {
			setting := models.Setting{
				MongoURL:     "",
				DatabaseName: "",
			}

			c.JSON(http.StatusOK, setting)
			return
		}

		urlString := fmt.Sprintf("%v", url)
		nmaeString := fmt.Sprintf("%v", nmae)

		setting := models.Setting{
			MongoURL:     urlString,
			DatabaseName: nmaeString,
		}

		c.JSON(http.StatusOK, setting)
	})

	routerV1.PUT("/config", func(c *gin.Context) {

		var setting models.Setting
		if err := c.ShouldBindJSON(&setting); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mutex.Lock()
		memoryMap["mongodbURL"] = setting.MongoURL
		memoryMap["mongodbName"] = setting.DatabaseName
		mutex.Unlock()

		c.JSON(http.StatusOK, setting)
	})

	routerV1.POST("/task", func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var task *models.Task

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mutex.RLock()
		url, existsa := memoryMap["mongodbURL"]
		nmae, existsb := memoryMap["mongodbName"]
		mutex.RUnlock()

		if !existsa {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !existsb {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		urlString := fmt.Sprintf("%v", url)
		nmaeString := fmt.Sprintf("%v", nmae)

		db, err := connectMongo(urlString, nmaeString)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collectionTask := db.Collection("task")

		// Insert the book into the MongoDB collection
		result, err := collectionTask.InsertOne(ctx, task)
		if err != nil {
			return
		}

		// Get the inserted book's ID from the result
		insertedID, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			return
		}

		// Query the database to get the complete book object with the generated ID
		var createdTask models.TaskResponse
		err = collectionTask.FindOne(ctx, bson.M{"_id": insertedID}).Decode(&createdTask)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, createdTask)
	})

	routerV1.GET("/task", func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		mutex.RLock()
		url, existsa := memoryMap["mongodbURL"]
		nmae, existsb := memoryMap["mongodbName"]
		mutex.RUnlock()

		if !existsa {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !existsb {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		urlString := fmt.Sprintf("%v", url)
		nmaeString := fmt.Sprintf("%v", nmae)

		db, err := connectMongo(urlString, nmaeString)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collectionTask := db.Collection("task")

		// Find all tasks in the MongoDB collection
		cursor, err := collectionTask.Find(ctx, bson.M{})
		if err != nil {
			return
		}
		defer cursor.Close(ctx)

		// Decode the results into a slice of tasks
		var tasks []models.TaskResponse
		for cursor.Next(ctx) {
			var task models.TaskResponse
			if err := cursor.Decode(&task); err != nil {
				return
			}
			tasks = append(tasks, task)
		}

		if tasks == nil {
			tasks = []models.TaskResponse{}
		}

		c.JSON(http.StatusOK, tasks)
	})

	return app
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
