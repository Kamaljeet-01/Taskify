package db

import (
	"context"
	"time"

	"github.com/harsh082ip/todolist-CLI-golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(title string, desc string) (string, error) {

	client, err := DbConnection()
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	objectID := primitive.NewObjectID()
	var todo models.TodoModel = models.TodoModel{
		Id:          objectID,
		Title:       title,
		Description: desc,
		Time:        time.Now().Format("2006-01-02 15:04:05"),
		IsCompleted: false,
	}

	coll := client.Database("Todos-App").Collection("todos")
	_, err = coll.InsertOne(context.TODO(), todo)
	if err != nil {
		return "", err
	}

	return "Todo Added Successfully", err
}
