package service

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/bells307/everydaypic/internal/domain/image/dto"
	"github.com/bells307/everydaypic/internal/domain/image/model"
)

// Интерфейс репозитория с информацией о картинках
type imageRepository interface {
	Add(ctx context.Context, name, fileName, userID string) (model.Image, error)
	Delete(ctx context.Context, imageID string) error
	Get(ctx context.Context, dto dto.GetImages) ([]model.Image, error)
}

// Интерфейс хранилища файлов
type imageFileStorage interface {
	Upload(ctx context.Context, name string, filename string, fileSize int64, data io.Reader) error
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
	// image := model.Image{
	// 	ID:       uuid.NewString(),
	// 	Name:     dto.Name,
	// 	FileName: dto.FileName,
	// 	Created:  time.Now(),
	// }

	image, err := s.imageRepository.Add(ctx, dto.Name, dto.FileName, dto.UserID)
	if err != nil {
		return image, fmt.Errorf("can't upload image to repository: %v", err)
	}

	if err := s.imageFileStorage.Upload(ctx, image.ID, image.FileName, dto.FileSize, dto.Data); err != nil {
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

// Получить изображения по фильтру
func (s *ImageService) Get(ctx context.Context, dto dto.GetImages) ([]model.Image, error) {
	return s.imageRepository.Get(ctx, dto)
}

// Получить ссылку на скачивание
func (s *ImageService) GetUrl(ctx context.Context, imageID string) (url.URL, error) {
	return s.imageFileStorage.GetUrl(ctx, imageID)
}
