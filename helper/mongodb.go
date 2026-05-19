package helper

import (
	"context"
	"time"

	"backend/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollection returns a MongoDB collection reference
func GetCollection(collectionName string) *mongo.Collection {
	return config.MongoClient.Database("kampus").Collection(collectionName)
}

// GetContext returns a context with timeout for MongoDB operations
func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
