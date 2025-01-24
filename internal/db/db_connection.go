package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnection() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")

	if len(uri) == 0 {
		// fmt.Println("MONGODB_URI not set")
		return nil, fmt.Errorf("MONGODB_URI not set")
	}
	// fmt.Println(uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Use defer to ensure Disconnect is called before returning
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		fmt.Println("Error disconnecting from MongoDB:", err)
	// 	}
	// }()
	return client, nil
}

// fmt.Println(time.Now())
// currentTime := time.Now()
// fmt.Println("Current Date and Time:", currentTime)

// // You can also format the output as needed
// fmt.Println("Formatted Date and Time:", currentTime.Format("2006-01-02 15:04:05"))
