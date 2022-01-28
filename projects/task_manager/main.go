package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Task - Model of a basic task
type Task struct {
	ID    primitive.ObjectID
	Title string
	Body  string
}

const (
	connectionTimeout = 5
	connectionURI     = "mongodb://127.0.0.1:27017/taskManager"
)

func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))

	if err != nil {
		log.Printf("Failed to create client %v ", client)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Printf("Failed to connect to cluster %v ", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Printf("Failed to ping to cluster %v ", err)
	}
	fmt.Println("Connected to Mongodb")
	return client, ctx, cancel
}

func Create(task *Task) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	task.ID = primitive.NewObjectID()
	result, err := client.Database("tasksManger").Collection("tasks").insertOne(ctx, task)

	if err != nil {
		log.Printf("Could not create task %v", err)
		return primitive.NilObjectID, err
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

// GET REQUEST HANDLER
func handleGetTasks(c *gin.Context) {
	var tasks []Task
	var task Task
	task.Title = "Bake some cake"
	task.Body = `- Make dough
	- Eat everything before baking
	- Pretend you didn't want to bake something in the first place`
	tasks = append(tasks, task)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

// PUT REQUEST HANDLER
func handleCreateTask(c *gin.Context) {
	var task Task

	if err := c.ShouldBind(&task); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	id, err := Create(&task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func main() {
	r := gin.Default()
	r.GET("/tasks/", handleGetTasks)
	r.PUT("/tasks/", handleCreateTask)
	r.Run()
}
