package usecase

import (
	"context"
	"fmt"
	"net/url"
	"time"

	fileService "github.com/bells307/everydaypic/internal/domain/file/service"
	fileServiceDTO "github.com/bells307/everydaypic/internal/domain/file/service/dto"
	"github.com/bells307/everydaypic/internal/domain/image/model"
	imageService "github.com/bells307/everydaypic/internal/domain/image/service"
	imageServiceDTO "github.com/bells307/everydaypic/internal/domain/image/service/dto"
	"github.com/bells307/everydaypic/internal/domain/image/usecase/dto"
	"github.com/google/uuid"
)

const IMAGE_BUCKET = "images"

type ImageUsecase struct {
	imageService *imageService.ImageService
	fileService  *fileService.FileService
}

func NewImageUsecase(imageService *imageService.ImageService, fileService *fileService.FileService) *ImageUsecase {
	return &ImageUsecase{imageService, fileService}
}

// Получить изображения по фильтру
func (u *ImageUsecase) GetImages(ctx context.Context, dto dto.GetImages) ([]model.Image, error) {
	return u.imageService.GetByFilter(ctx, imageServiceDTO.GetImagesFilter{
		ID:       dto.ID,
		FileName: dto.FileName,
	})
}

// Проверить существование изображения по ID
func (u *ImageUsecase) CheckImageExists(ctx context.Context, imageID string) (bool, error) {
	return u.imageService.CheckExists(ctx, imageID)
}

// Добавить изображение
func (u *ImageUsecase) CreateImage(ctx context.Context, dto dto.CreateImage) (img model.Image, err error) {
	img = model.Image{
		ID:       uuid.NewString(),
		Name:     dto.Name,
		FileName: dto.FileName,
		UserID:   dto.UserID,
		Created:  time.Now(),
	}

	// TODO: предусмотреть откат изменений при неудаче одного из этапов (транзакция?)

	// Загружаем в image service
	if err := u.imageService.Create(ctx, img); err != nil {
		return img, fmt.Errorf("can't create image %s (%s): %v", img.Name, img.ID, err)
	}

	// Загружаем в file service
	fileServiceDTO := fileServiceDTO.UploadFile{
		Bucket:   IMAGE_BUCKET,
		Name:     img.ID,
		Filename: dto.FileName,
		FileSize: dto.FileSize,
		Data:     dto.Data,
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
