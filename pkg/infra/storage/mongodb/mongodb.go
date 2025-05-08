package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type DBConnector interface {
	GetCollection(name string) *mongo.Collection
	Close() error
}

func New(uri string, dbName string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017"))
	if err != nil {
		return nil, fmt.Errorf("❌ failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("❌ failed to ping MongoDB: %w", err)
	}

	db := client.Database(dbName)

	fmt.Println("✅ MongoDB connected")

	return &Database{
		Client:   client,
		Database: db,
	}, nil
}

func (d *Database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return d.Client.Disconnect(ctx)
}

func (d *Database) GetCollection(name string) *mongo.Collection {
	return d.Database.Collection(name)
}
