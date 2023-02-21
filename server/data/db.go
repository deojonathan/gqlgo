package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *DBClient

type DBClient struct {
	client *mongo.Client
}

func NewDBClient() *DBClient {
	return &DBClient{}
}

func (dbc *DBClient) GetClient() *mongo.Client {
	return dbc.client
}

func (dbc *DBClient) GetDatabase(name string) *mongo.Database {
	return dbc.client.Database(name)
}

func (dbc *DBClient) GetCollection(dbName, collName string) *mongo.Collection {
	return dbc.client.Database(dbName).Collection(collName)
}

func (dbc *DBClient) Connect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://db:27017"))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		return err
	}

	dbc.client = client
	return nil
}

func (dbc *DBClient) Disconnect() (err error) {
	if dbc.client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = dbc.client.Disconnect(ctx); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
		return err
	}
	return nil
}
