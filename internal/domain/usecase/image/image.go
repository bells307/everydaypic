package image_usecase

import (
	"context"

	"github.com/bells307/everydaypic/internal/domain/service"
)

type imageUsecase struct {
	imageService service.ImageService
}

func NewImageUsecase(imageService service.ImageService) *imageUsecase {
	return &imageUsecase{imageService: imageService}
}

func (u *imageUsecase) CreateImage(ctx context.Context, dto CreateImageDTO) error {
	
	return nil
}
