package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoModel struct {
	Id          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"desc,omitempty"`
	Time        string             `bson:"time,omitempty"`
	IsCompleted bool               `bson:"iscompleted"`
}
