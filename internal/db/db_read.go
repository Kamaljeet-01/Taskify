package db

import (
	"context"

	"github.com/harsh082ip/todolist-CLI-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListAllTodos() ([]models.TodoModel, error) {

	client, err := DbConnection()
	if err != nil {
		return nil, err
	}

	coll := client.Database("Todos-App").Collection("todos")

	filter := bson.D{}

	res, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer res.Close(context.TODO())

	var todos []models.TodoModel

	for res.Next(context.TODO()) {
		var todo models.TodoModel
		if err := res.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	// Check for errors during cursor iteration
	if err := res.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func ListTodoByID(id string) (models.TodoModel, error) {

	client, err := DbConnection()
	if err != nil {
		return models.TodoModel{}, err
	}

	coll := client.Database("Todos-App").Collection("todos")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.TodoModel{}, err
	}

	// Creating a filter to fetch by ID
	filter := bson.D{{"_id", objectID}}

	var todo models.TodoModel

	err = coll.FindOne(context.TODO(), filter).Decode(&todo)
	if err != nil {
		return models.TodoModel{}, err
	}

	return todo, nil
}

func ListTodoByStatus(status bool) ([]models.TodoModel, error) {

	client, err := DbConnection()
	if err != nil {
		return nil, err
	}

	coll := client.Database("Todos-App").Collection("todos")

	filter := bson.D{{"iscompleted", status}}

	var todos []models.TodoModel

	res, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer res.Close(context.TODO())

	for res.Next(context.TODO()) {
		var todo models.TodoModel
		if err := res.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	// Check for errors during cursor iteration
	if err := res.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
