package usecase

import (
	"context"
	"fmt"
	"net/url"

	"github.com/bells307/everydaypic/internal/dto"
	"github.com/bells307/everydaypic/internal/model"
	"github.com/bells307/everydaypic/internal/service"
)

const IMAGE_BUCKET = "images"

type ImageUsecase struct {
	imageService *service.ImageService
	fileService  *service.FileService
}

func NewImageUsecase(imageService *service.ImageService, fileService *service.FileService) *ImageUsecase {
	return &ImageUsecase{imageService, fileService}
}

// Получить изображения по фильтру
func (u *ImageUsecase) GetImages(ctx context.Context, getImages dto.GetImages) ([]model.Image, error) {
	return u.imageService.GetByFilter(ctx, getImages.ID, getImages.FileName)
}

// Проверить существование изображения по ID
func (u *ImageUsecase) CheckImageExists(ctx context.Context, imageID string) (bool, error) {
	return u.imageService.CheckExists(ctx, imageID)
}

// Добавить изображение
func (u *ImageUsecase) CreateImage(ctx context.Context, createImage dto.CreateImage) (img model.Image, err error) {
	// TODO: предусмотреть откат изменений при неудаче одного из этапов (транзакция?)

	img, err = u.imageService.Create(ctx, createImage.Name, createImage.FileName, createImage.UserID)
	if err != nil {
		return img, fmt.Errorf("can't create image %s (%s): %v", img.Name, img.ID, err)
	}

	// Загружаем в file service
	fileServiceDTO := dto.UploadFile{
		Bucket:   IMAGE_BUCKET,
		Name:     img.ID,
		Filename: createImage.FileName,
		FileSize: createImage.FileSize,
		Data:     createImage.Data,
	}

	if err := u.fileService.Upload(ctx, fileServiceDTO); err != nil {
		return img, fmt.Errorf("can't upload image %s (%s) to file storage: %v", img.Name, img.ID, err)
	}

	return
}

// Удалить изображение
func (u *ImageUsecase) DeleteImage(ctx context.Context, imageID string) error {
	if err := u.imageService.Delete(ctx, imageID); err != nil {
		return fmt.Errorf("error deleting image %s from image repository: %v", imageID, err)
	}

	if err := u.fileService.Delete(ctx, IMAGE_BUCKET, imageID); err != nil {
		return fmt.Errorf("error deleting image %s from file storage: %v", imageID, err)
	}

	return nil
}

// Получить ссылку на скачивание
func (u *ImageUsecase) GetDownloadUrl(ctx context.Context, imageID string) (*url.URL, error) {
	return u.fileService.GetUrl(ctx, IMAGE_BUCKET, imageID)
}
