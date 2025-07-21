package data

import (
	"context"
	"errors"

	"task_manager_with_mongo/db"
	"task_manager_with_mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task

	cur, err := db.TaskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cur.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskByID(id string) (*models.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	var task models.Task
	err = db.TaskCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func CreateTask(task models.Task) (*models.Task, error) {
	task.ID = primitive.NewObjectID()
	_, err := db.TaskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
		},
	}

	result, err := db.TaskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("task not found")
	}

	updatedTask.ID = objectID
	return &updatedTask, nil
}

func DeleteTask(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	res, err := db.TaskCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
