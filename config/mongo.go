package config

import (
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/hariszaki17/go-api-clean/exception"
)
// NewMongoDatabase expose global
func NewMongoDatabase(configuration Config) *mongo.Database  {
	ctx, cancel := NewMongoContext()
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(configuration.Get("MONGO_URI")))
	exception.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exception.PanicIfNeeded(err)

	database := client.Database(configuration.Get("MONGO_DATABASE"))
	return database
}

// NewMongoContext expose global
func NewMongoContext() (context.Context, context.CancelFunc)  {
	return context.WithTimeout(context.Background(), 10*time.Second)
}