package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoRead(ctx *context.Context) (*mongo.Database, error) {
	mongoUsername := os.Getenv("MONGO_DB_USERNAME")
	mongoPassword := os.Getenv("MONGO_DB_PASSWORD")
	mongoServer := os.Getenv("MONGO_DB_SERVER")
	mongoDatabaseName := os.Getenv("MONGO_DB_NAME")
	fullUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", mongoUsername, mongoPassword, mongoServer, mongoDatabaseName)
	client, err := ConnectMongo(fullUrl, *ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(mongoDatabaseName), nil
}

func MongoWrite(ctx *context.Context) (*mongo.Database, error) {
	mongoUsername := os.Getenv("MONGO_DB_USERNAME")
	mongoPassword := os.Getenv("MONGO_DB_PASSWORD")
	mongoServer := os.Getenv("MONGO_DB_SERVER")
	mongoDatabaseName := os.Getenv("MONGO_DB_NAME")
	fullUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", mongoUsername, mongoPassword, mongoServer, mongoDatabaseName)
	fmt.Println(fullUrl)
	client, err := ConnectMongo(fullUrl, *ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(mongoDatabaseName), nil

}

func ConnectMongo(url string, ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client()
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(10)

	client, err := mongo.NewClient(clientOptions.ApplyURI(url))

	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CloseConnection(client *mongo.Client, ctx *context.Context) {
	client.Disconnect(*ctx)
}
