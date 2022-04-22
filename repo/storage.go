package repo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mongoEntity struct {
	ID     string              `bson:"_id"`
	Entity bson.D              `bson:"entity"`
	CAS    uint64              `bson:"cas"`
	TTL    *primitive.DateTime `bson:"ttl,omitempty"`
}
