package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoDB struct {
	db *mongo.Database
}
