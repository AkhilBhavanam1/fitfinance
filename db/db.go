package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHelperClient struct {
	ConnString string
	Client     *mongo.Client
}

type DB interface {
	ConnectToDB(string) error
}

func (dbclient *DBHelperClient) ConnectToDB() error {

	clientOptions := options.Client().ApplyURI(dbclient.ConnString)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	dbclient.Client = client
	log.Println("Connected to MongoDB!")
	return nil
}
