package usecase

import (
	"context"

	"github.com/bells307/everydaypic/internal/domain/entity"
	"github.com/bells307/everydaypic/internal/domain/service"
)

type ImageUsecase interface {
	CreateImage(ctx context.Context, name, filename string, data []byte) (entity.Image, error)
}

type imageUsecase struct {
	imageService service.ImageService
}

func NewImageUsecase(imageService service.ImageService) *imageUsecase {
	return &imageUsecase{imageService: imageService}
}

func (u *imageUsecase) CreateImage(ctx context.Context, name, filename string, data []byte) (entity.Image, error) {
	return u.imageService.CreateImage(ctx, name, filename, data)
}
