package usecase

import (
	"context"
	"errors"

	"github.com/bells307/everydaypic/internal/domain/dto"
)

var ErrNotFound = errors.New("file not found")

type ImageService interface {
	CreateImage(ctx context.Context, dto dto.CreateImage) (string, error)
	DeleteImage(ctx context.Context, id string) error
	DownloadImage(ctx context.Context, id string) ([]byte, error)
}

type imageUsecase struct {
	imageService ImageService
}

func NewImageUsecase(imageService ImageService) *imageUsecase {
	return &imageUsecase{imageService}
}

func (u *imageUsecase) CreateImage(ctx context.Context, dto dto.CreateImage) (string, error) {
	return u.imageService.CreateImage(ctx, dto)
}

func (u *imageUsecase) DeleteImage(ctx context.Context, id string) error {
	return u.imageService.DeleteImage(ctx, id)
}

func (u *imageUsecase) DownloadImage(ctx context.Context, id string) ([]byte, error) {
	return u.imageService.DownloadImage(ctx, id)
}
