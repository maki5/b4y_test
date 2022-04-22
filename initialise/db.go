package initialise

import (
	"context"
	"github.com/maki5/b4y_test/repo"
	"github.com/maki5/b4y_test/repo/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName   = "b4y_test"
	maxCount = 0
)

func MongoClient() (*mongo.Client, error) {
	ctx := context.Background()
	return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}

func CacheRepo(client *mongo.Client) *repo.CacheRepo {
	mongoClient := storage.NewMongoClient(client, dbName)
	return repo.NewCacheRepo(mongoClient, dbName, maxCount)
}
