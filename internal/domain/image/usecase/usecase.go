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
	Get(ctx context.Context, dto dto.GetImages) ([]model.Image, error)
	GetUrl(ctx context.Context, imageID string) (url.URL, error)
}

type ImageUsecase struct {
	imageService imageService
}

func NewImageUsecase(imageService imageService) *ImageUsecase {
	return &ImageUsecase{imageService}
}

// Получить изображения по фильтру
func (u *ImageUsecase) GetImages(ctx context.Context, dto dto.GetImages) ([]model.Image, error) {
	return u.imageService.Get(ctx, dto)
}

// Добавить изображение
func (u *ImageUsecase) AddImage(ctx context.Context, dto dto.CreateImage) (model.Image, error) {
	image, err := u.imageService.Add(ctx, dto)
	if err != nil {
		return model.Image{}, fmt.Errorf("can't add image: %v", err)
	}

	return image, nil
}

// Удалить изображение
func (u *ImageUsecase) DeleteImage(ctx context.Context, imageID string) error {
	return u.imageService.Delete(ctx, imageID)
}

// Получить ссылку на скачивание
func (u *ImageUsecase) GetDownloadUrl(ctx context.Context, imageID string) (url.URL, error) {
	return u.imageService.GetUrl(ctx, imageID)
}
