package repo

import (
	"context"
	"encoding/json"
	"github.com/maki5/b4y_test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"strconv"
	"time"
)

type MongoClient interface {
	Create(ctx context.Context, collection string, entity interface{}) error
	GetOne(ctx context.Context, collection string, filter bson.M) ([]byte, error)
	Delete(ctx context.Context, collection string, filter bson.M, opts ...*options.FindOneAndDeleteOptions) error
	Update(ctx context.Context, collection string, filter bson.M, entity interface{}) error
	Count(ctx context.Context, collection string) (int64, error)
}

type CacheRepo struct {
	client         MongoClient
	collectionName string
	maxCount       int
}

func NewCacheRepo(client MongoClient, collectionName string, maxCount int) *CacheRepo {
	return &CacheRepo{client: client, collectionName: collectionName, maxCount: maxCount}
}

func (c *CacheRepo) GetCacheByKey(ctx context.Context, key string) (*models.Cache, error) {
	data, err := c.client.GetOne(ctx, c.collectionName,
		bson.M{"entity.key": key,
			"entity.ttl": bson.M{
				"$lt": time.Now().Unix(),
			},
		},
	)

	if err != nil {
		return nil, err
	}

	cache := &models.Cache{}
	err = json.Unmarshal(data, cache)
	if err != nil {
		return nil, err
	}

	updatedCache := *cache
	updatedCache.TTL = time.Now().Add(time.Duration(rand.Int31n(1000000000))).Unix()
	err = c.client.Update(ctx, c.collectionName, bson.M{"entity.key": key}, updatedCache)
	if err != nil {
		return nil, err
	}

	return cache, nil
}

func (c *CacheRepo) CreateCache(ctx context.Context, key string) error {
	cache := models.Cache{
		Key:   key,
		Value: strconv.Itoa(int(time.Now().Unix())),
		TTL:   time.Now().Add(time.Duration(rand.Int31n(1000000000))).Unix(),
	}

	count, err := c.client.Count(ctx, c.collectionName)
	if err != nil {
		return err
	}

	if int(count) >= c.maxCount {
		err := c.deleteFirstCache(ctx)
		if err != nil {
			return err
		}
	}

	return c.client.Create(ctx, c.collectionName, cache)
}

func (c *CacheRepo) deleteFirstCache(ctx context.Context) error {
	filer := bson.M{"$natural": -1}
	options := options.FindOneAndDelete()
	options.SetSort(filer)
	return c.client.Delete(ctx, c.collectionName, bson.M{}, options)
}
