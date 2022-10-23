package usecase

import (
	"context"

	"github.com/bells307/everydaypic/internal/domain/entity"
	"github.com/bells307/everydaypic/internal/domain/service"
)

type ImageUsecase interface {
	UploadImage(ctx context.Context, img entity.Image, data []byte) error
	DeleteImage(ctx context.Context, id string) error
}

type imageUsecase struct {
	imageService service.ImageService
}

func NewImageUsecase(imageService service.ImageService) *imageUsecase {
	return &imageUsecase{imageService: imageService}
}

func (u *imageUsecase) UploadImage(ctx context.Context, img entity.Image, data []byte) error {
	return u.imageService.CreateImage(ctx, img, data)
}

func (u *imageUsecase) DeleteImage(ctx context.Context, id string) error {
	return u.imageService.DeleteImage(ctx, id)
}
