package common

import "go.mongodb.org/mongo-driver/mongo"

func NoDocuments(err error) bool {
	if err == mongo.ErrNoDocuments {
		return true
	}

	return false
}
