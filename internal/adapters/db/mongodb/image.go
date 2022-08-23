package mongodb

import (
	"context"
	"log"

	"github.com/bells307/everydaypic/internal/domain/entity"
	"github.com/bells307/everydaypic/pkg/mongodb"
)

type imageStorage struct {
	db *mongodb.MongoDB
}

func NewImageStorage(db *mongodb.MongoDB) *imageStorage {
	return &imageStorage{db: db}
}

func (s *imageStorage) CreateImage(ctx context.Context, name, filename string, data []byte) (entity.Image, error) {
	log.Println("creating image in mongo ...")
	return entity.Image{}, nil
}

func (s *imageStorage) DeleteImage(ctx context.Context, id string) error {
	log.Printf("deleting image %s in mongo ...\n", id)
	return nil
}
