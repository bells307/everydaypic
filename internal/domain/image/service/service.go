package service

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/bells307/everydaypic/internal/domain/image/dto"
	"github.com/bells307/everydaypic/internal/domain/image/model"
	"github.com/google/uuid"
)

// Интерфейс репозитория с информацией о картинках
type imageRepository interface {
	Add(ctx context.Context, image model.Image) error
	Delete(ctx context.Context, imageID string) error
	GetByID(ctx context.Context, imageID string) (image model.Image, err error)
}

// Интерфейс хранилища файлов
type imageFileStorage interface {
	Upload(ctx context.Context, name string, filename string) error
	Delete(ctx context.Context, name string) error
	GetUrl(ctx context.Context, imageID string) (url.URL, error)
}

type ImageService struct {
	imageRepository  imageRepository
	imageFileStorage imageFileStorage
}

func NewImageService(imageRepository imageRepository, imageFileStorage imageFileStorage) *ImageService {
	return &ImageService{imageRepository, imageFileStorage}
}

// Добавить изображение
func (s *ImageService) Add(ctx context.Context, dto dto.CreateImage) (model.Image, error) {
	image := model.Image{
		ID:       uuid.NewString(),
		FileName: dto.FileName,
		Created:  time.Time{},
	}

	if err := s.imageRepository.Add(ctx, image); err != nil {
		return image, fmt.Errorf("can't upload image to repository: %v", err)
	}

	if err := s.imageFileStorage.Upload(ctx, image.ID, image.FileName); err != nil {
		return image, fmt.Errorf("can't upload image to file storage: %v", err)
	}

	return image, nil
}

// Удалить изображение
func (s *ImageService) Delete(ctx context.Context, imageID string) error {
	if err := s.imageRepository.Delete(ctx, imageID); err != nil {
		return fmt.Errorf("can't delete image from repository: %v", err)
	}

	// // Удаляем из хранилища
	if err := s.imageFileStorage.Delete(ctx, imageID); err != nil {
		return fmt.Errorf("can't delete image from file storage: %v", err)
	}

	return nil
}

// Получить изображение по ID
func (s *ImageService) GetByID(ctx context.Context, imageID string) (image model.Image, err error) {
	return s.imageRepository.GetByID(ctx, imageID)
}

// Получить ссылку на скачивание
func (s *ImageService) GetUrl(ctx context.Context, imageID string) (url.URL, error) {
	return s.imageFileStorage.GetUrl(ctx, imageID)
}
