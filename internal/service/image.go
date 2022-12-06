package service

import (
	"context"
	"time"

	"github.com/bells307/everydaypic/internal/model"
	"github.com/google/uuid"
)

// Интерфейс репозитория с информацией о картинках
type imageRepository interface {
	Create(ctx context.Context, image model.Image) error
	Delete(ctx context.Context, imageID string) error
	GetByFilter(ctx context.Context, IDs, fileNames []string) ([]model.Image, error)
	CheckExists(ctx context.Context, imageID string) (bool, error)
}

type ImageService struct {
	repo imageRepository
}

func NewImageService(imageRepository imageRepository) *ImageService {
	return &ImageService{imageRepository}
}

// Добавить изображение
func (s *ImageService) Create(ctx context.Context, name, fileName, userID string) (model.Image, error) {
	image := model.Image{
		ID:       uuid.NewString(),
		Name:     name,
		FileName: fileName,
		UserID:   userID,
		Created:  time.Now(),
	}

	return image, s.repo.Create(ctx, image)
}

// Удалить изображение
func (s *ImageService) Delete(ctx context.Context, imageID string) error {
	return s.repo.Delete(ctx, imageID)
}

// Получить изображения по фильтру
func (s *ImageService) GetByFilter(ctx context.Context, IDs, fileNames []string) ([]model.Image, error) {
	return s.repo.GetByFilter(ctx, IDs, fileNames)
}

// Проверить существование изображения
func (s *ImageService) CheckExists(ctx context.Context, imageID string) (bool, error) {
	return s.repo.CheckExists(ctx, imageID)
}
