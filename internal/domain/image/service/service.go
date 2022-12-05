package service

import (
	"context"

	"github.com/bells307/everydaypic/internal/domain/image/model"
	repoDTO "github.com/bells307/everydaypic/internal/domain/image/repository/dto"
	"github.com/bells307/everydaypic/internal/domain/image/service/dto"
)

// Интерфейс репозитория с информацией о картинках
type imageRepository interface {
	Create(ctx context.Context, image model.Image) error
	Delete(ctx context.Context, imageID string) error
	GetByFilter(ctx context.Context, dto repoDTO.GetImagesFilter) ([]model.Image, error)
	CheckExists(ctx context.Context, imageID string) (bool, error)
}

type ImageService struct {
	repo imageRepository
}

func NewImageService(imageRepository imageRepository) *ImageService {
	return &ImageService{imageRepository}
}

// Добавить изображение
func (s *ImageService) Create(ctx context.Context, image model.Image) error {
	return s.repo.Create(ctx, image)
}

// Удалить изображение
func (s *ImageService) Delete(ctx context.Context, imageID string) error {
	return s.repo.Delete(ctx, imageID)
}

// Получить изображения по фильтру
func (s *ImageService) GetByFilter(ctx context.Context, dto dto.GetImagesFilter) ([]model.Image, error) {
	return s.repo.GetByFilter(ctx, dto.ToRepoDTO())
}

func (s *ImageService) CheckExists(ctx context.Context, imageID string) (bool, error) {
	return s.repo.CheckExists(ctx, imageID)
}
