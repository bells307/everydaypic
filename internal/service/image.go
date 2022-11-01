package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bells307/everydaypic/internal/dto"
	"github.com/bells307/everydaypic/internal/entity"
)

// Ошибки
// Изображение не найдено
var ErrNotFound = errors.New("image not found")

// Интерфейс хранилища изображений
type ImageStorage interface {
	// Получить изображения
	GetImages(ctx context.Context, dto dto.GetImages) ([]entity.Image, error)
	// Создать изображение
	CreateImage(ctx context.Context, dto dto.CreateImage) (entity.Image, error)
	// Удалить изображение
	DeleteImage(ctx context.Context, id string) error
	// Скачать изображение
	DownloadImage(ctx context.Context, id string) ([]byte, error)
	// Создать изображение дня
	CreateImageOfTheDay(ctx context.Context, img entity.DayImage) error
}

// Сервис работы с изображениями
type imageService struct {
	storage ImageStorage
}

func NewImageService(storage ImageStorage) *imageService {
	return &imageService{storage: storage}
}

// Получить изображения
func (s *imageService) GetImages(ctx context.Context, dto dto.GetImages) ([]entity.Image, error) {
	return s.storage.GetImages(ctx, dto)
}

// Создать изображение
func (s *imageService) CreateImage(ctx context.Context, dto dto.CreateImage) (entity.Image, error) {
	return s.storage.CreateImage(ctx, dto)
}

// Удалить изображение
func (s *imageService) DeleteImage(ctx context.Context, id string) error {
	return s.storage.DeleteImage(ctx, id)
}

// Скачать изображение
func (s *imageService) DownloadImage(ctx context.Context, id string) ([]byte, error) {
	return s.storage.DownloadImage(ctx, id)
}

// Установить изображение дня
func (s *imageService) SetImageOfTheDay(ctx context.Context, id string) error {
	imgs, err := s.storage.GetImages(ctx, dto.GetImages{
		ID:       []string{id},
		FileName: []string{},
	})

	if err != nil {
		return fmt.Errorf("error getting images: %v", err)
	}

	img := imgs[0]

	dayImg := entity.DayImage{
		Image: img,
		SetAt: time.Now(),
	}

	return s.storage.CreateImageOfTheDay(ctx, dayImg)
}
