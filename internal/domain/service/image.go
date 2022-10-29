package service

import (
	"context"

	"github.com/bells307/everydaypic/internal/domain/dto"
	"github.com/bells307/everydaypic/internal/domain/entity"
)

type ImageStorage interface {
	GetImages(ctx context.Context, dto dto.GetImages) ([]entity.Image, error)
	CreateImage(ctx context.Context, dto dto.CreateImage) (entity.Image, error)
	DeleteImage(ctx context.Context, id string) error
	DownloadImage(ctx context.Context, id string) ([]byte, error)
}

type imageService struct {
	storage ImageStorage
}

func NewImageService(storage ImageStorage) *imageService {
	return &imageService{storage: storage}
}

func (s *imageService) GetImages(ctx context.Context, dto dto.GetImages) ([]entity.Image, error) {
	return s.storage.GetImages(ctx, dto)
}

func (s *imageService) CreateImage(ctx context.Context, dto dto.CreateImage) (entity.Image, error) {
	return s.storage.CreateImage(ctx, dto)
}

func (s *imageService) DeleteImage(ctx context.Context, id string) error {
	return s.storage.DeleteImage(ctx, id)
}

func (s *imageService) DownloadImage(ctx context.Context, id string) ([]byte, error) {
	return s.storage.DownloadImage(ctx, id)
}
