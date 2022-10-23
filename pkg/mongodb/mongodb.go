package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	Uri    string
	DbName string
}

type MongoDB struct {
	db *mongo.Database
}

func NewMongoDB(cfg MongoDBConfig) (*MongoDB, error) {
	opts := options.Client().ApplyURI(cfg.Uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, fmt.Errorf("can't connect to mongodb: %v", err)
	}

	db := client.Database(cfg.DbName)
	if db == nil {
		return nil, fmt.Errorf("can't create a handle to mongodb database %s", cfg.DbName)
	}

	return &MongoDB{
		db,
	}, nil
}

func (m *MongoDB) UploadFile(filename string, data []byte) error {
	bucket, err := gridfs.NewBucket(m.db)
	if err != nil {
		return fmt.Errorf("error creating gridfs bucket: %v", err)
	}

	uploadStream, err := bucket.OpenUploadStream(filename)
	if err != nil {
		return fmt.Errorf("error opening upload stream for %s: %v", filename, err)
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to upload stream for %s: %v", filename, err)
	}

	return nil
}
