package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTodoById(id string) (string, error) {

	client, err := DbConnection()
	if err != nil {
		return "", err
	}
	// fetch collection
	coll := client.Database("Todos-App").Collection("todos")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	// creating a filter
	filter := bson.D{{"_id", objectId}}

	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return "", err
	}

	// fmt.Println(res)
	if res.DeletedCount != 0 {
		return fmt.Sprintf("Successfully Deleted Todo %v \n", id), nil
	} else {
		return fmt.Sprintf("Todo %v is already Deleted", id), nil
	}
}

func DeleteAllTodos() (string, error) {

	client, err := DbConnection()
	if err != nil {
		return "", err
	}
	// fetch collection
	coll := client.Database("Todos-App").Collection("todos")

	// include all docs
	filter := bson.D{}

	res, err := coll.DeleteOne(context.TODO(), filter)

	if res.DeletedCount != 0 {
		return fmt.Sprintf("Successfully Deleted All Todos"), nil
	} else {
		return fmt.Sprintf("Nothing To Delete"), nil
	}

}

func DeleteCompOrIncompTODOS(todoStatus bool) (string, error) {

	client, err := DbConnection()
	if err != nil {
		return "", err
	}

	// fetch collection
	coll := client.Database("Todos-App").Collection("todos")

	// include docs accordingly
	var filter primitive.D
	if todoStatus {
		filter = bson.D{{"iscompleted", todoStatus}}
	} else {
		filter = bson.D{{"iscompleted", todoStatus}}
	}

	res, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return "", err
	}

	if todoStatus {
		if res.DeletedCount != 0 {
			return fmt.Sprintf("Successfully Deleted all Completed Todos"), nil
		} else {
			return fmt.Sprintf("There is no Complete Todo to delete"), nil
		}
	} else {
		if res.DeletedCount != 0 {
			return fmt.Sprintf("Successfully Deleted all Completed Todos"), nil
		} else {
			return fmt.Sprintf("There is no Incomplete Todo to delete"), nil
		}
	}
}
