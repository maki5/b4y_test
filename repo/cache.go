package repo

import (
	"context"
	"encoding/json"
	"github.com/maki5/b4y_test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

type CacheRepo struct {
	client         *mongo.Client
	dbName         string
	collectionName string
}

func NewCacheRepo(client *mongo.Client, dbName string, collectionName string) *CacheRepo {
	return &CacheRepo{client: client, dbName: dbName, collectionName: collectionName}
}

func (c *CacheRepo) GetCacheByKey(ctx context.Context, key string) (*models.Cache, error) {
	db := c.client.Database(c.dbName).Collection(c.collectionName)

	entity := mongoEntity{}
	err := db.FindOne(ctx, bson.M{"entity.key": key}).Decode(&entity)
	if err != nil {
		return nil, err
	}

	data, err := bson.MarshalExtJSON(entity.Entity, false, false)
	if err != nil {
		return nil, err
	}

	cache := &models.Cache{}
	err = json.Unmarshal(data, cache)
	if err != nil {
		return nil, err
	}

	return cache, nil
}

func (c *CacheRepo) CreateCache(ctx context.Context, key string) error {
	cache := models.Cache{
		Key:   key,
		Value: strconv.Itoa(int(time.Now().Unix())),
	}

	filter := bson.M{
		"entity": cache,
	}

	_, err := c.client.Database(c.dbName).Collection(c.collectionName).InsertOne(ctx, filter, nil)
	if err != nil {
		return err
	}

	return nil
}
