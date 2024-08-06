package db

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbUri  = "mongodb://localhost:27017"
	dbName = "excel_api"
)

const (
	CollectionStudnets = "students"
)

var (
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(dbUri).SetAuth(options.Credential{
			Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
			Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		})

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}

func GetMongoDatabase() *mongo.Database {
	client, err := GetMongoClient()
	if err != nil {
		return nil
	}

	database := client.Database(dbName)

	return database
}

func Seed() {
	err := SeedStudentsCollection()
	if err != nil {
		log.Fatal(err)
	}
}
