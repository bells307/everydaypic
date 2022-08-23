package service

import (
	"context"

	"github.com/bells307/everydaypic/internal/domain/entity"
)

type ImageStorage interface {
	CreateImage(ctx context.Context, name, filename string, data []byte) (entity.Image, error)
	DeleteImage(ctx context.Context, id string) error
}

type ImageService interface {
	CreateImage(ctx context.Context, name, filename string, data []byte) (entity.Image, error)
	DeleteImage(ctx context.Context, id string) error
}

type imageService struct {
	storage ImageStorage
}

func NewImageService(storage ImageStorage) *imageService {
	return &imageService{storage: storage}
}

func (s *imageService) CreateImage(ctx context.Context, name, filename string, data []byte) (entity.Image, error) {
	return s.storage.CreateImage(ctx, name, filename, data)
}

func (s *imageService) DeleteImage(ctx context.Context, id string) error {
	return s.storage.DeleteImage(ctx, id)
}
