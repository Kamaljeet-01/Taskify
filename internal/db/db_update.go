package db

import (
	"context"
	"fmt"
	"os"

	// "github.com/harsh082ip/todolist-CLI-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTodo(id, title, desc string) (string, error) {
	client, err := DbConnection()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	coll := client.Database("Todos-App").Collection("todos")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	filter := bson.D{{"_id", objectID}}

	update := bson.D{{"$set", bson.D{{"title", title}, {"desc", desc}}}}
	res, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)
		os.Exit(1)
	}

	// fmt.Println(res.ModifiedCount)
	// fmt.Println(filter)
	// fmt.Println(update)
	if res.MatchedCount == 0 {
		return fmt.Sprintf("No todos matched with the Id"), nil
	} else if res.ModifiedCount == 0 {
		return fmt.Sprintf("Todo is already have the same data"), nil
	} else if res.ModifiedCount != 0 {
		return fmt.Sprintf("Todo Updated Successfully"), nil
	} else {
		return fmt.Sprintf("Some Error Occured"), nil
	}

}

func MarkTodosAsCompOrIncomp(id string, isCompleted bool, markAll bool) (string, error) {

	client, err := DbConnection()
	if err != nil {
		return "", err
	}

	coll := client.Database("Todos-App").Collection("todos")
	var filter primitive.D
	if markAll {
		filter = bson.D{}
	} else {
		ObjectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return "", err
		}
		filter = bson.D{{"_id", ObjectID}}
	}

	var update primitive.D
	if isCompleted {
		update = bson.D{{"$set", bson.D{{"iscompleted", true}}}}
	} else {
		update = bson.D{{"$set", bson.D{{"iscompleted", false}}}}
	}

	if !markAll {
		res, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return "", err
		}

		if res.MatchedCount == 0 {
			return fmt.Sprintf("No todos matched with the Id"), nil
		} else if res.ModifiedCount == 0 {
			if isCompleted {
				return fmt.Sprintf("Todo is already completed"), nil
			} else {
				return fmt.Sprintf("Todo is already incompleted"), nil
			}
		} else if res.ModifiedCount != 0 {
			return fmt.Sprintf("Todo Updated Successfully"), nil
		} else {
			return fmt.Sprintf("Some Error Occured"), nil
		}
	}

	res, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return "", nil
	}

	if res.MatchedCount == 0 {
		return fmt.Sprintf("No todos matched with the Id"), nil
	} else if res.ModifiedCount == 0 {
		if isCompleted {
			return fmt.Sprintf("All Todos are already completed"), nil
		} else {
			return fmt.Sprintf("All Todos are already incompleted"), nil
		}
	} else if res.ModifiedCount != 0 {
		return fmt.Sprintf("Todo Updated Successfully"), nil
	} else {
		return fmt.Sprintf("Some Error Occured"), nil
	}
}
