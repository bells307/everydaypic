package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type imageStorage struct {
	db *mongo.Database
}

func NewImageStorage(db *mongo.Database) *imageStorage {
	return &imageStorage{db: db}
}
