package service

import (
	"context"
	"io"
	"net/url"

	"github.com/bells307/everydaypic/internal/dto"
)

// Интерфейс хранилища файлов
type fileRepository interface {
	// Загрузить файл
	Upload(ctx context.Context, name, fileName, bucket string, fileSize int64, data io.ReadSeeker) error
	// Удалить файл
	Delete(ctx context.Context, bucket, name string) error
	// Получить URL файла
	GetUrl(ctx context.Context, bucket, name string) (*url.URL, error)
}

// Сервис работы с загружаемыми файлами
type FileService struct {
	repo fileRepository
}

func NewFileService(repo fileRepository) *FileService {
	return &FileService{repo}
}

func (s *FileService) Upload(ctx context.Context, uploadFile dto.UploadFile) error {
	return s.repo.Upload(
		ctx, uploadFile.Name,
		uploadFile.Filename,
		uploadFile.Bucket,
		uploadFile.FileSize,
		uploadFile.Data,
	)
}

func (s *FileService) Delete(ctx context.Context, bucket, name string) error {
	return s.repo.Delete(ctx, bucket, name)
}

func (s *FileService) GetUrl(ctx context.Context, bucket, name string) (*url.URL, error) {
	return s.repo.GetUrl(ctx, bucket, name)
}
