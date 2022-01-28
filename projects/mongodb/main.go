package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	// ctx will be used to create a deadline  for process
	// deadline will be 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// mongo.Connect return mongo.Client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connected successfully")
	return nil
}

// this method closes mongodb connection and cancel context
func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	// cancel context
	defer cancel()

	// client.Disconnect method also has deadline
	// returns error if any
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func insertOne(client *mongo.Client, ctx context.Context, database string, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	collection := client.Database(database).Collection(col)

	result, err := collection.InsertOne(ctx, doc)
	return result, err

}

// func insertMany(client *mongo.Client, ctx context.Context, database string, col string, docs interface{}) (*mongo.InsertManyResult, error) {

// 	collection := client.Database(database).Collection(col)

// 	result, err := collection.InsertMany(ctx, docs)
// 	return result, err

// }

func main() {
	client, ctx, cancel, err := connect("mongodb://127.0.0.1:27017/taskManager")

	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	// ping(client, ctx)
	var document interface{}

	document = bson.D{
		{"rollNo", 175},
		{"maths", 80},
		{"science", 90},
		{"computer", 95},
	}

	insertOneResult, err := insertOne(client, ctx, "taskManager", "tasks", document)
	if err != nil {
		panic(err)
	}

	fmt.Println(insertOneResult)
}
