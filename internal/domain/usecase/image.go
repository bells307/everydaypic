package usecase

import (
	"context"
	"errors"

	"github.com/bells307/everydaypic/internal/domain/dto"
	"github.com/bells307/everydaypic/internal/domain/entity"
)

var ErrNotFound = errors.New("file not found")

type ImageService interface {
	GetImages(ctx context.Context, dto dto.GetImages) ([]entity.Image, error)
	CreateImage(ctx context.Context, dto dto.CreateImage) (entity.Image, error)
	DeleteImage(ctx context.Context, id string) error
	DownloadImage(ctx context.Context, id string) ([]byte, error)
}

type imageUsecase struct {
	imageService ImageService
}

func NewImageUsecase(imageService ImageService) *imageUsecase {
	return &imageUsecase{imageService}
}

func (u *imageUsecase) GetImages(ctx context.Context, dto dto.GetImages) ([]entity.Image, error) {
	return u.imageService.GetImages(ctx, dto)
}

func (u *imageUsecase) CreateImage(ctx context.Context, dto dto.CreateImage) (entity.Image, error) {
	return u.imageService.CreateImage(ctx, dto)
}

func (u *imageUsecase) DeleteImage(ctx context.Context, id string) error {
	return u.imageService.DeleteImage(ctx, id)
}

func (u *imageUsecase) DownloadImage(ctx context.Context, id string) ([]byte, error) {
	return u.imageService.DownloadImage(ctx, id)
}
