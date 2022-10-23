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
	return &imageStorage{db}
}

func (s *imageStorage) CreateImage(ctx context.Context, img entity.Image, data []byte) error {
	log.Println("creating image in mongo ...")
	s.db.UploadFile(img.FileName, data)
	return nil
}

func (s *imageStorage) DeleteImage(ctx context.Context, id string) error {
	log.Printf("deleting image %s in mongo ...\n", id)
	return nil
}
