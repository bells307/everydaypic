package usecase

import (
	"context"
	"fmt"
	"net/url"

	"github.com/bells307/everydaypic/internal/domain/image/dto"
	"github.com/bells307/everydaypic/internal/domain/image/model"
)

// Интерфейс сервиса работы с изображениями
type imageService interface {
	Add(ctx context.Context, dto dto.CreateImage) (model.Image, error)
	Delete(ctx context.Context, imageID string) error
	GetByID(ctx context.Context, imageID string) (model.Image, error)
	GetUrl(ctx context.Context, imageID string) (url.URL, error)
}

type imageUsecase struct {
	imageService imageService
}

func NewImageUsecase(imageService imageService) *imageUsecase {
	return &imageUsecase{imageService}
}

// Получить изображение по ID
func (u *imageUsecase) GetByID(ctx context.Context, imageID string) (model.Image, error) {
	return u.imageService.GetByID(ctx, imageID)
}

// Добавить изображение
func (u *imageUsecase) AddImage(ctx context.Context, dto dto.CreateImage) (model.Image, error) {
	image, err := u.imageService.Add(ctx, dto)
	if err != nil {
		return model.Image{}, fmt.Errorf("can't add image: %v", err)
	}

	return image, nil
}

// Удалить изображение
func (u *imageUsecase) DeleteImage(ctx context.Context, imageID string) error {
	return u.imageService.Delete(ctx, imageID)
}

// Получить ссылку на скачивание
func (u *imageUsecase) GetDownloadUrl(ctx context.Context, imageID string) (url.URL, error) {
	return u.imageService.GetUrl(ctx, imageID)
}
