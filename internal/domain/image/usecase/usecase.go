package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/bells307/everydaypic/internal/domain/image/dto"
	"github.com/bells307/everydaypic/internal/domain/image/model"
)

// Ошибки
// Изображение не найдено
var ErrNotFound = errors.New("image not found")

// Интерфейс сервиса работы с изображениями
type imageService interface {
	Add(ctx context.Context, dto dto.CreateImage) (model.Image, error)
	Delete(ctx context.Context, imageID string) error
	Get(ctx context.Context, dto dto.GetImages) ([]model.Image, error)
	GetUrl(ctx context.Context, imageID string) (*url.URL, error)
	CheckExists(ctx context.Context, imageID string) (bool, error)
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

func (u *ImageUsecase) CheckImageExists(ctx context.Context, imageID string) (bool, error) {
	return u.imageService.CheckExists(ctx, imageID)
}

// Получить изображение по ID
func (u *ImageUsecase) GetImageByID(ctx context.Context, imageID string) (model.Image, error) {
	imgs, err := u.imageService.Get(ctx, dto.GetImages{
		ID:       []string{imageID},
		FileName: []string{},
	})

	if err != nil {
		return model.Image{}, fmt.Errorf("error while getting image by ID: %v", err)
	}

	if len(imgs) == 0 {
		return model.Image{}, ErrNotFound
	} else {
		return imgs[0], nil
	}
}

// Добавить изображение
func (u *ImageUsecase) AddImage(ctx context.Context, dto dto.CreateImage) (model.Image, error) {
	return u.imageService.Add(ctx, dto)
}

// Удалить изображение
func (u *ImageUsecase) DeleteImage(ctx context.Context, imageID string) error {
	return u.imageService.Delete(ctx, imageID)
}

// Получить ссылку на скачивание
func (u *ImageUsecase) GetDownloadUrl(ctx context.Context, imageID string) (*url.URL, error) {
	return u.imageService.GetUrl(ctx, imageID)
}
