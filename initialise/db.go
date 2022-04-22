package initialise

import (
	"context"
	"github.com/maki5/b4y_test/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName         = "b4y_test"
	collectionName = "caches"
)

func MongoClient() (*mongo.Client, error) {
	ctx := context.Background()
	return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}

func CacheRepo(mongoClient *mongo.Client) *repo.CacheRepo {
	return repo.NewCacheRepo(mongoClient, dbName, collectionName)
}
