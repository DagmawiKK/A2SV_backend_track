package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"task_manager_with_mongo/db"
	"task_manager_with_mongo/router"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("mongo_url")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	db.TaskCollection = client.Database("taskmanager").Collection("tasks")

	r := router.SetupRouter()
	r.Run(":8080")
}
