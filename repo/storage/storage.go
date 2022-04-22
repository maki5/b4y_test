package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongoEntity struct {
	ID     string              `bson:"_id"`
	Entity bson.D              `bson:"entity"`
	CAS    uint64              `bson:"cas"`
	TTL    *primitive.DateTime `bson:"ttl,omitempty"`
}

type MongoClient struct {
	client *mongo.Client
	dbName string
}

func NewMongoClient(client *mongo.Client, dbName string) *MongoClient {
	return &MongoClient{
		client: client,
		dbName: dbName,
	}
}

func (c *MongoClient) Create(ctx context.Context, collection string, entity interface{}) error {
	filter := bson.M{
		"entity": entity,
	}

	_, err := c.client.Database(c.dbName).Collection(collection).InsertOne(ctx, filter, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *MongoClient) GetOne(ctx context.Context, collection string, filter bson.M) ([]byte, error) {
	db := c.client.Database(c.dbName).Collection(collection)

	entity := mongoEntity{}
	err := db.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}

	return bson.MarshalExtJSON(entity.Entity, false, false)
}

func (c *MongoClient) Delete(ctx context.Context, collection string, filter bson.M, opts ...*options.FindOneAndDeleteOptions) error {
	res := c.client.Database(c.dbName).Collection(collection).FindOneAndDelete(ctx, filter, opts...)
	return res.Err()
}

func (c *MongoClient) Update(ctx context.Context, collection string, filter bson.M, entity interface{}) error {
	query := bson.D{{Key: "$set", Value: bson.D{{Key: "entity", Value: entity}}}}

	res, err := c.client.Database(c.dbName).Collection(collection).UpdateOne(ctx, filter, query)
	if err != nil {
		return err
	}

	log.Printf("updated %v entities", res.ModifiedCount)

	return nil
}

func (c *MongoClient) Count(ctx context.Context, collection string) (int64, error) {
	return c.client.Database(c.dbName).Collection(collection).EstimatedDocumentCount(ctx, nil)
}
