package mongodb

import (
	"context"
	"log"

	"github.com/bells307/everydaypic/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type imageStorage struct {
	db *mongo.Database
}

func NewImageStorage(db *mongo.Database) *imageStorage {
	return &imageStorage{db: db}
}

func (s *imageStorage) CreateImage(ctx context.Context, name, filename string, data []byte) (entity.Image, error) {
	log.Println("creating image in mongo ...")
	return entity.Image{}, nil
}
